package logic

import (
	"{{.gomod}}/types"

	"github.com/gin-gonic/gin"
)

type {{.handler}} struct {
}

func (h *{{.handler}}) Req() interface{} {
	return &types.{{.request}}{}
}
func (h *{{.handler}}) HttpMethod() string {
	return "{{.method}}"
}

func (h *{{.handler}}) Do(iReq interface{}, ctx *gin.Context) (rsp interface{}, err error) {
	req := iReq.(*types.{{.request}})
	return req,nil //delete it!
}
