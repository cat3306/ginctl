package gogen

import (
	_ "embed"
	"strings"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed middleware.tpl
var middlewareImplementCode string

func genMiddleware(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	middlewares := getMiddleware(api)
	for _, m := range middlewares {
		err := genFile(fileGenConfig{
			dir:             routerDir,
			subdir:          middlewareDir,
			filename:        strings.ToLower(m) + ".go",
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
