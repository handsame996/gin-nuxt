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

	g := models.CreateGlobal(gormDB.InitializeGorm(c.Mysql), configs.InitZap(c.Zap))

	initHttp.Server(g, c.Http)
}
