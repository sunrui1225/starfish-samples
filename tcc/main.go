package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gotrx/starfish/pkg/client"
	"github.com/gotrx/starfish/pkg/client/config"
	"github.com/gotrx/starfish/pkg/client/tcc"
	"github.com/gotrx/starfish/pkg/client/tm"
	"github.com/gotrx/starfish/samples/tcc/service"
)

func main() {
	r := gin.Default()

	config.InitConfWithDefault("testService")
	client.NewRpcClient()
	tcc.InitTCCResourceManager()

	tm.Implement(service.ProxySvc)
	tcc.ImplementTCC(service.TccProxyServiceA)
	tcc.ImplementTCC(service.TccProxyServiceB)
	tcc.ImplementTCC(service.TccProxyServiceC)

	r.GET("/commit", func(c *gin.Context) {
		service.ProxySvc.TCCCommitted(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/rollback", func(c *gin.Context) {
		service.ProxySvc.TCCCanceled(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
