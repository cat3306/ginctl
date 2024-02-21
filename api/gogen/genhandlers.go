package gogen

import (
	_ "embed"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed handler.tpl
var handlerTemplate string

//go:embed rsp.tpl
var rspTemplate string

func genHandlers(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	err := genFile(fileGenConfig{
		dir:             dir,
		subdir:          handlerDir,
		filename:        "handler" + ".go",
		templateName:    "handlerTemplate",
		category:        category,
		templateFile:    handlerTemplateFile,
		builtinTemplate: handlerTemplate,
		data: map[string]string{
			"gomod": rootPkg,
		},
	})
	if err != nil {
		return err
	}
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          handlerDir,
		filename:        "rsp" + ".go",
		templateName:    "rspTemplate",
		category:        category,
		templateFile:    rspTemplateFile,
		builtinTemplate: rspTemplate,
		data:            map[string]string{},
	})
}
