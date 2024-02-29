package gogen

import (
	_ "embed"
	"path"
	"strings"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/config"
)

//go:embed logic.tpl
var logicTemplate string

func genLogic(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {

	// raw, _ := json.Marshal(api.Service)
	// fmt.Printf("%+v\n", string(raw))
	var newFiles []string
	for _, group := range api.Service.Groups {
		for _, route := range group.Routes {
			newFiles = append(newFiles, path.Join(dir, logicDir, strings.ToLower(route.Handler)+".go"))
		}
	}
	delNotExistFiles(newFiles, path.Join(dir, logicDir))
	for _, group := range api.Service.Groups {
		for _, route := range group.Routes {
			reqTypeName := route.RequestTypeName()
			rspTypeName := route.ResponseTypeName()
			err := genFile(fileGenConfig{
				dir:             dir,
				subdir:          logicDir,
				filename:        strings.ToLower(route.Handler) + ".go",
				templateName:    "logicTemplate",
				category:        category,
				templateFile:    logicTemplateFile,
				builtinTemplate: logicTemplate,
				data: map[string]interface{}{
					"gomod":           rootPkg,
					"handler":         StrFirstLetterUp(route.Handler),
					"method":          strings.ToUpper(route.Method),
					"request":         StrFirstLetterUp(route.RequestTypeName()),
					"isRequestEmpty":  reqTypeName == "",
					"isResponseEmpty": rspTypeName == "",
					"reqAndRsp":       reqTypeName == "" && rspTypeName == "",
					"response":        StrFirstLetterUp(route.ResponseTypeName()),
				},
			})
			if err != nil {
				return nil
			}
		}
	}
	return nil
}
