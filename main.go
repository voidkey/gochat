package main

import (
	"gochat/router"
	"gochat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8080")
}
