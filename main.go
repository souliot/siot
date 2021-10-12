package main

import (
	logs "github.com/souliot/siot-log"
)

var (
	appname = "siot"
	version = "1.0.0"
	desc    = "A cli for generate go project"
)

func main() {
	initLog()
	app := &App{
		Name:    appname,
		Usage:   desc,
		Version: version,
	}
	err := app.run()
	if err != nil {
		logs.Error("程序运行错误：", err)
	}
}
