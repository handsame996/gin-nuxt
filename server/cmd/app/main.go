package main

import (
	"example/templalte/internal/configs"
	"example/templalte/internal/db/gormDB"
	"example/templalte/internal/httpServer/initHttp"
	"example/templalte/internal/models"
)

func main() {
	var c configs.ConfigModel
	configs.InitializeViper(&c)

	var g models.Global
	g.DB = gormDB.InitializeGorm(c.Mysql)

	initHttp.InitHttpServer(&g, c.Http)
}
