package logic

import (
	{{if .reqAndRsp}}
	{{else}}"{{.gomod}}/types"{{end}}

	"github.com/gin-gonic/gin"
)

type {{.handler}} struct {
}

func (h *{{.handler}}) Request() interface{} {
	{{if .isRequestEmpty}}return nil
	{{else}}return &types.{{.request}}{}{{end}}
}

func (h *{{.handler}}) Response() interface{} {
	{{if .isResponseEmpty}}return nil
	{{else}}return &types.{{.response}}{}{{end}}
}

func (h *{{.handler}}) HttpMethod() string {
	return "{{.method}}"
}

func (h *{{.handler}}) Do(iReq interface{}, ctx *gin.Context) (rsp interface{}, err error) {
	{{if .isRequestEmpty}}
	{{else}}req := iReq.(*types.{{.request}})
	_ = req //delete it !
	{{end}}
	return 
}
