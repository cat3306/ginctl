package gogen

import (
	_ "embed"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed util.tpl
var utilTemplate string

func genUtil(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          utilDir,
		filename:        "client.go",
		templateName:    "utilTemplate",
		category:        category,
		templateFile:    utilTemplateFile,
		builtinTemplate: utilTemplate,
		data:            map[string]any{},
	})
}
