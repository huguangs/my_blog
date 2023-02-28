package main

import (
	"go_blog/model"
	"go_blog/routes"
	"go_blog/utils"
)

func main() {
	model.InitDb()
	r := routes.InitRouter()
	r.Run(utils.HttpPort)

}
