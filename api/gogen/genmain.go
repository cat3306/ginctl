package gogen

import (
	_ "embed"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed main.tpl
var mainTemplate string

//go:embed gomod.tpl
var gomodTemlate string

func genMain(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec, component string) error {

	err := genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        "main" + ".go",
		templateName:    "mainTemplate",
		category:        category,
		templateFile:    mainTemplateFile,
		builtinTemplate: mainTemplate,
		data: map[string]any{
			"gomod":        rootPkg,
			"hasComponent": component != "",
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
			"goversion": getGoVersion(),
		},
	})
}

func getGoVersion() string {
	cmd := exec.Command("go", "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	v := string(out)
	vlist := strings.Split(v, " ")
	version := ``
	if len(vlist) >= 3 {
		version = vlist[2]
	}
	if len(version) >= 2 {
		version = version[2:]
	}
	fmt.Println(version)
	return version
}
