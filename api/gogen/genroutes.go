package gogen

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed router.tpl
var routerTemlate string

//go:embed customrouter.tpl
var customRouterTemlate string

//go:embed autorouter.tpl
var autoRouterTemlate string

func genRoutes(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {

	routerSrc := genRouterSrc(api)
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

	autogenFileName := "autogenrouter" + ".go"
	middlewareList := getMiddleware(api)
	filename := path.Join(dir, routerDir, autogenFileName)
	os.Remove(filename)
	err = genFile(fileGenConfig{
		dir:             dir,
		subdir:          routerDir,
		filename:        autogenFileName,
		templateName:    "autoRouterTemlate",
		category:        category,
		templateFile:    autoRouterTemplateFile,
		builtinTemplate: autoRouterTemlate,
		data: map[string]any{
			"gomod":      rootPkg,
			"routersrc":  routerSrc,
			"middleware": len(middlewareList) != 0,
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

func genRouterSrc(api *spec.ApiSpec) string {

	srcTmplate := `%s.%s("%s",handler.GinWrapper(new(logic.%s)))`
	groupTmplate := `
	%s := engine.Group("%s"%s)
	{
		%s
	}
	`

	finalSrc := ``
	for _, group := range api.Service.Groups {
		gName := group.GetAnnotation("group")
		gName = handlerGroupStr(gName)
		prefix := group.GetAnnotation("prefix")
		midd := group.GetAnnotation("middleware")
		gPrefix := "engine"
		src := ``
		if gName != "" && prefix != "" {
			gPrefix = gName
		}
		for _, route := range group.Routes {
			src += fmt.Sprintf(srcTmplate, gPrefix, strings.ToUpper(route.Method), route.Path, StrFirstLetterUp(route.Handler)) + "\n\n"
		}
		if gName != "" && prefix != "" {
			midSrc := joinMiddleware(midd)
			src = fmt.Sprintf(groupTmplate, gName, prefix, midSrc, src)
		}
		finalSrc += src
	}
	return finalSrc
}
func joinMiddleware(s string) string {
	src := ``
	ms := strings.Split(s, ",")
	for _, v := range ms {
		src += fmt.Sprintf(`,middleware.%s()`, v)
	}
	return src
}
func handlerGroupStr(s string) string {
	if len(s) <= 1 {
		return s
	}
	s = strings.ReplaceAll(s, "/", "")
	return strings.ToLower(s[:1]) + s[1:]
}
