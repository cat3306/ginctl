package gogen

import (
	_ "embed"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
	"github.com/cat3306/ginctl/util/format"
)

const (
	configFile = "config"
)

//go:embed config.tpl
var configTemplate string

//go:embed configtype.tpl
var configTypeTemplate string

func genConfig(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, configFile)
	if err != nil {
		return err
	}
	err = genFile(fileGenConfig{
		dir:             dir,
		subdir:          configDir,
		filename:        filename + ".go",
		templateName:    "configTemplate",
		category:        category,
		templateFile:    configTemplateFile,
		builtinTemplate: configTemplate,
		data:            map[string]string{},
	})
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          configDir,
		filename:        filename + "type" + ".go",
		templateName:    "configTypeTemplate",
		category:        category,
		templateFile:    configTypeTemplateFile,
		builtinTemplate: configTypeTemplate,
		data:            map[string]string{},
	})
}
