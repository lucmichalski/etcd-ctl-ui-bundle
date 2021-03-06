package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

// Binding from JSON
type User struct {
	User      string   `form:"user" json:"user" xml:"user"`
	Password  string   `form:"password" json:"password" xml:"password"`
	Role      string   `form:"role" json:"role" xml:"role"`
	Endpoints []string `form:"endpoints" json:"endpoints" xml:"endpoints"`
}

func (handler *httpHanlder) login(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := NewEtcdClient(json.User, json.Password, json.Endpoints)
	if err == rpctypes.ErrGRPCAuthNotEnabled {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	tokenString, err := authMiddleware.MakeCredential(json.User, json.Password, json.Endpoints)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
