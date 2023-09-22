package main

import (
	"context"
	"gochat/router"
	"gochat/utils"
)

func main() {
	ctx := context.Background()
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis(ctx)

	r := router.Router()
	r.Run(":8080")
}
