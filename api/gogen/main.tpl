package main

import (
	"{{.gomod}}/applog"
	"{{.gomod}}/config"
	"{{.gomod}}/router"
	"flag"

)

func main() {
	var file string
	flag.StringVar(&file, "f", "", "use -f to bind config file")
	flag.Parse()
	config.Init(file)
	applog.Init()
	router.StartGinServer()
}
