package gogen

import (
	_ "embed"
	"path"
	"strings"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed middleware.tpl
var middlewareImplementCode string

var customMiddlewareImplement = `
package custom
// middleware custom 
`
var customRoutelogImplement = `
package custom

import (
	"bytes"
	"{{.gomod}}/applog"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomGinLogger(headerKey []string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		rawQuery := ctx.Request.URL.RawQuery
		var err error
		var body []byte
		if ctx.Request.Method == http.MethodPost {
			contentType := ctx.Request.Header.Get("Content-Type")
			if strings.Contains(contentType, "application/json") {
				body, err = io.ReadAll(ctx.Request.Body)
				if err != nil && !errors.Is(err, io.EOF) {
					applog.Logger.Sugar().Errorf("CustomGinLogger err:%s", err.Error())
					return
				}
				if body != nil {
					ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
				}

			}
		}
		ctx.Next()

		now := time.Now()
		latency := now.Sub(start)

		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		bodyStr := ""
		if body != nil {
			buffer := new(bytes.Buffer)
			if err := json.Compact(buffer, body); err != nil {
				fmt.Println(err)
			}
			bodyStr = buffer.String()
		}
		printHeader := ""
		for i, key := range headerKey {
			format := "%s:%s,"
			if i == len(headerKey)-1 {
				format = "%s:%s"
			}
			printHeader += fmt.Sprintf(format, key, ctx.Request.Header.Get(key))
		}
		//bodySize := ctx.Writer.Size()
		applog.Logger.Sugar().Infof("[GIN] [%d] [%v] [%s] [%s] [%s] [body:%s] [query:%s] [header:%s]",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			bodyStr,
			rawQuery,
			printHeader,
		)
	}
}

`

func genCustomMiddleware(rootPkg string) error {
	err := genFile(fileGenConfig{
		dir:             path.Join(routerDir, middlewareDir),
		subdir:          "custom",
		filename:        "readme" + ".go",
		templateName:    "customMiddlewareImplement",
		category:        category,
		templateFile:    "",
		builtinTemplate: customMiddlewareImplement,
		data:            map[string]string{},
	})
	if err != nil {
		return err
	}
	return genFile(fileGenConfig{
		dir:             path.Join(routerDir, middlewareDir),
		subdir:          "custom",
		filename:        "customlog" + ".go",
		templateName:    "customRoutelogImplement",
		category:        category,
		templateFile:    "",
		builtinTemplate: customRoutelogImplement,
		data: map[string]string{
			"gomod": rootPkg,
		},
	})
}
func genMiddleware(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	middlewares := getMiddleware(api)
	newFiles := make([]string, 0)
	for _, m := range middlewares {
		fileName := strings.ToLower(m) + ".go"
		newFiles = append(newFiles, path.Join(dir, middlewareDir, fileName))
	}
	delNotExistFiles(newFiles, path.Join(routerDir, middlewareDir))
	for _, m := range middlewares {
		fileName := strings.ToLower(m) + ".go"
		err := genFile(fileGenConfig{
			dir:             routerDir,
			subdir:          middlewareDir,
			filename:        fileName,
			templateName:    "middlewareImplementCode",
			category:        category,
			templateFile:    middlewareImplementCodeFile,
			builtinTemplate: middlewareImplementCode,
			data: map[string]interface{}{
				"middleware": m,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
