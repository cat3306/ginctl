package gogen

var customLogic = `
package custom
// logic custom 
`

func genCustomlogic() error {
	return genFile(fileGenConfig{
		dir:             logicDir,
		subdir:          "custom",
		filename:        "readme" + ".go",
		templateName:    "customLogic",
		category:        category,
		templateFile:    "",
		builtinTemplate: customLogic,
		data:            map[string]string{},
	})
}
