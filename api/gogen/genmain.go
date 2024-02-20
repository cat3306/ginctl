package gogen

import (
	_ "embed"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed main.tpl
var mainTemplate string

//go:embed gomod.tpl
var gomodTemlate string

func genMain(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {

	err := genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        "main" + ".go",
		templateName:    "mainTemplate",
		category:        category,
		templateFile:    mainTemplateFile,
		builtinTemplate: mainTemplate,
		data: map[string]string{
			"gomod": rootPkg,
		},
	})
	if err != nil {
		return err
	}
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        "go" + ".mod",
		templateName:    "gomodTemlate",
		category:        category,
		templateFile:    gomodTemplateFile,
		builtinTemplate: gomodTemlate,
		data: map[string]string{
			"gomod":     rootPkg,
			"goversion": "1.21.5",
		},
	})
}
