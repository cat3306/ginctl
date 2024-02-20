package gogen

import (
	_ "embed"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed applog.tpl
var appLogTemplate string

func genAppLog(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          applogDir,
		filename:        "log" + ".go",
		templateName:    "appLogTemplate",
		category:        category,
		templateFile:    appLogTemplateFile,
		builtinTemplate: appLogTemplate,
		data:            map[string]string{},
	})
}
