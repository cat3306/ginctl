package main

import (
	"{{.gomod}}/applog"
	"{{.gomod}}/config"
	"{{.gomod}}/router"
	{{if .hasComponent}}"{{.gomod}}/component"{{end}}
	"flag"

)

func main() {
	var file string
	flag.StringVar(&file, "f", "", "use -f to bind config file")
	flag.Parse()
	config.Init(file)
	applog.Init()
	{{if .hasComponent}}
	err := component.Init()
	if err != nil {
		panic(err)
	}
	{{end}}
	router.StartGinServer()
}
