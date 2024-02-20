package gogen

import (
	_ "embed"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed router.tpl
var routerTemlate string

//go:embed customrouter.tpl
var customRouterTemlate string

func genRoutes(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	err := genFile(fileGenConfig{
		dir:             dir,
		subdir:          routerDir,
		filename:        "router" + ".go",
		templateName:    "routerTemlate",
		category:        category,
		templateFile:    routerTemplateFile,
		builtinTemplate: routerTemlate,
		data: map[string]string{
			"gomod": rootPkg,
		},
	})
	if err != nil {
		return err
	}
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          routerDir,
		filename:        "customrouter" + ".go",
		templateName:    "customRouterTemlate",
		category:        category,
		templateFile:    customRouterTemplateFile,
		builtinTemplate: customRouterTemlate,
		data:            map[string]string{},
	})
}
