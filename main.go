package main

import (
	"curdapp/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tools.ParseConfig("./config/config.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()
	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}
