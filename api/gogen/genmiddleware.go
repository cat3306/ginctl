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

func genCustomMiddleware() error {
	return genFile(fileGenConfig{
		dir:             path.Join(routerDir, middlewareDir),
		subdir:          "custom",
		filename:        "readme" + ".go",
		templateName:    "customMiddlewareImplement",
		category:        category,
		templateFile:    "",
		builtinTemplate: customMiddlewareImplement,
		data:            map[string]string{},
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
