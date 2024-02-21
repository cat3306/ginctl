package router

import (
	"{{.gomod}}/handler"
	"{{.gomod}}/logic"

	"github.com/gin-gonic/gin"
)

func AutoRegister(engine *gin.Engine) {
	{{.routersrc}}
}