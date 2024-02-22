package router

import (
	"{{.gomod}}/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartGinServer() {
	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	CustomRegister(engine)
	AutoRegister(engine)
	err := engine.Run(fmt.Sprintf("%s:%s", config.AppConf.Host, config.AppConf.Port))
	if err != nil {
		panic(err)
	}
}