package router

import (
	"{{.gomod}}/config"
	"{{.gomod}}/handler"
	"{{.gomod}}/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartGinServer() {
	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	AutoRegister(engine)
	CustomRegister(engine)
	err := engine.Run(fmt.Sprintf("%s:%s", config.AppConf.Host, config.AppConf.Port))
	if err != nil {
		panic(err)
	}
}

func AutoRegister(engine *gin.Engine) {
	
}
