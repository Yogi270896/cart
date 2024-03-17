package main

import (
	"cart/internal/app"
	"cart/internal/confi"
)

var (
	appConf *confi.AppConfig
)

func init() {
	appConf = confi.NewConfig()
}

func main() {
	app.StartApplication(appConf)
}
