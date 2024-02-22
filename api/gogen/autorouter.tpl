package router

import (
	"{{.gomod}}/handler"
	"{{.gomod}}/logic"
	
	{{if .middleware}}"{{.gomod}}/router/middleware"{{end}}
	"github.com/gin-gonic/gin"
)

func AutoRegister(engine *gin.Engine) {
	{{.routersrc}}
}