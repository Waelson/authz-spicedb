package authz

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const Schema = `
definition blog/user {}

definition blog/post {
	relation reader: blog/user
	relation writer: blog/user

	permission read = reader + writer
	permission write = writer
}
`

type Client interface {
	CheckPermission(ctx context.Context) (bool, error)
	StoreRelationship(ctx context.Context, relationships []Relationship) error
	ApplySchema(schema string) error
}

type client struct {
	authzedClient *authzed.Client
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

func (a *client) CheckPermission(ctx context.Context) (bool, error) {
	return true, nil
}

func (a *client) ApplySchema(schema string) error {
	request := &pb.WriteSchemaRequest{Schema: schema}
	_, err := a.authzedClient.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
		return err
	}
	return nil
}

func (a *client) StoreRelationship(ctx context.Context, relationships []Relationship) error {

	return nil
}

func NewAuthZClient(host string, token string) (Client, error) {
	c, err := authzed.NewClient(
		host,
		grpcutil.WithInsecureBearerToken(token),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return &client{}, err
	}

	return &client{
		authzedClient: c,
	}, nil
}
