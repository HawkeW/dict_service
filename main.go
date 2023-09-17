package main

import (
	"dictService/global"
	"dictService/models"
	"dictService/routers"
)

func main() {
	global.InitConfig()
	models.InitDb()
	routers.StartServer()
}
