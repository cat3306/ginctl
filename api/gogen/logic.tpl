package logic

import (
	"{{.gomod}}/types"

	"github.com/gin-gonic/gin"
)

type {{.handler}} struct {
}

func (h *{{.handler}}) Request() interface{} {
	{{if .isTypeEmpty}}return nil
	{{else}}return &types.{{.request}}{}{{end}}
}

func (h *{{.handler}}) Response() interface{} {
	return &types.{{.response}}{}
}

func (h *{{.handler}}) HttpMethod() string {
	return "{{.method}}"
}

func (h *{{.handler}}) Do(iReq interface{}, ctx *gin.Context) (rsp interface{}, err error) {
	{{if .isTypeEmpty}}
	{{else}}req := iReq.(*types.{{.request}})
	_ = req //delete it !
	{{end}}
	return 
}
