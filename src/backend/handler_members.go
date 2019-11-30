package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/clientv3"
)

func (handler *httpHanlder) ClusterMembers(c *gin.Context) {
	//cli, err := NewEtcdClient(c.MustGet("user.name").(string), c.MustGet("user.password").(string))
	//if err != nil {
	//c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	//return
	//}
	cli := c.MustGet("etcdClient").(*clientv3.Client)
	defer cli.Close()
	resp, err := cli.MemberList(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//memberjson, _ := json.Marshal(resp.Members)
	c.JSON(http.StatusOK, gin.H{"members": resp.Members})
}
