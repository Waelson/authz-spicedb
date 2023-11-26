package controller

import (
	"github.com/authz-spicedb/internal/authz"
	"github.com/gin-gonic/gin"
)

type AuthzController interface {
	StoreRelationship(c *gin.Context)
}

type authzController struct {
	authzCliente authz.Client
}

func (a *authzController) StoreRelationship(c *gin.Context) {

}

func NewAuthzController(client authz.Client) AuthzController {
	return &authzController{
		authzCliente: client,
	}
}
