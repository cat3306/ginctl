package gogen

import (
	_ "embed"
	"strconv"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

const (
	defaultPort = 8888
)

//go:embed etc.tpl
var etcTemplate string

func genEtc(dir string, cfg *config.Config, api *spec.ApiSpec, component string) error {
	service := api.Service
	host := "0.0.0.0"
	port := strconv.Itoa(defaultPort)
	data := genComponentDataMap(component)
	data["serviceName"] = service.Name
	data["host"] = host
	data["port"] = port
	data["release"] = true
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        "config.yaml",
		templateName:    "etcTemplate",
		category:        category,
		templateFile:    etcTemplateFile,
		builtinTemplate: etcTemplate,
		data:            data,
	})
}
