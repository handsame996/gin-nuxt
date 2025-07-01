package main

import (
	"example/template/internal/configs"
	"example/template/internal/db/gormDB"
	"example/template/internal/httpServer/initHttp"
	"example/template/internal/models"
)

func main() {
	var c configs.ConfigModel
	configs.InitializeViper(&c)

	var g models.Global
	g.DB = gormDB.InitializeGorm(c.Mysql)
	g.Logger = configs.InitZap(c.Zap)

	initHttp.InitHttpServer(&g, c.Http)
}
