package handler
import (
	"{{.gomod}}/applog"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Req() interface{}
	Do(interface{}) (interface{}, error)
	HttpMethod() string
}

func GinWrapper(h Handler) gin.HandlerFunc {
	fName := reflect.Indirect(reflect.ValueOf(h)).Type().Name()
	return func(c *gin.Context) {
		req := h.Req()
		var err error
		if req != nil {
			if h.HttpMethod() == http.MethodPost {
				err = c.BindJSON(req)
			} else if h.HttpMethod() == http.MethodGet {
				err = c.BindQuery(req)
			}
		}
		if err != nil {
			applog.Logger.Sugar().Errorf("%s params invalid,err:%s", fName, err.Error())
			return
		}

		rsp, err := h.Do(req)
		if err != nil {
			applog.Logger.Sugar().Errorf("%s err:%s,req:%+v", fName, err.Error(), req)
			RspError(c, err.Error())
			return
		}
		RspOk(c, rsp)
	}
}
