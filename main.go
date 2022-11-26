package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	router.Router()

	//r := router.Router()
	//r.Run(":8081")
}
