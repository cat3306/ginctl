package gogen

import (
	_ "embed"
	"os"
	"path"
	"strings"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed middleware.tpl
var middlewareImplementCode string

func genMiddleware(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	middlewares := getMiddleware(api)
	for _, v := range middlewares {

	}

	rmFiles()
	for _, m := range middlewares {
		fileName := strings.ToLower(m) + ".go"
		filename := path.Join(dir, middlewareDir, fileName)
		os.Remove(filename)
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
