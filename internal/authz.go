package internal

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type AuthZClient interface {
	CheckPermission(ctx context.Context) (bool, error)
	StoreRelationship(ctx context.Context, relationships []Relationship) error
	ApplySchema(schema string) error
}

type Relationship struct {
	Resource Resource
	Relation string
	Subject  Subject
}

type Resource struct {
	Type string
	Id   string
}

type Subject struct {
	Type string
	Id   string
}

type authZClient struct {
	authZedClient *authzed.Client
}

func (a *authZClient) CheckPermission(ctx context.Context) (bool, error) {
	return true, nil
}

func (a *authZClient) ApplySchema(schema string) error {
	request := &pb.WriteSchemaRequest{Schema: schema}
	_, err := a.authZedClient.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
		return err
	}
	return nil
}

func (a *authZClient) StoreRelationship(ctx context.Context, relationships []Relationship) error {
	return nil
}

func NewAuthZClient(host string, token string) (AuthZClient, error) {
	client, err := authzed.NewClient(
		host,
		grpcutil.WithInsecureBearerToken(token),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return &authZClient{}, err
	}

	return &authZClient{
		authZedClient: client,
	}, nil
}
