package main

import (
	"example/template/internal/auth/jwt"
	"example/template/internal/configs"
	"example/template/internal/db/gormDB"
	"example/template/internal/httpServer/initHttp"
	"example/template/internal/models"
)

func main() {
	var c configs.ConfigModel
	configs.InitializeViper(&c)

	g := models.CreateGlobal(gormDB.InitializeGorm(c.Mysql), configs.InitZap(c.Zap), jwt.NewJWT(c.JWT))

	initHttp.Server(g, c.Http)
}
