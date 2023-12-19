package main

import (
	"version1_copy/model"
	"version1_copy/router"
	"version1_copy/utils"
)

func main() {
	utils.Init()
	model.InitDB()
	r := router.SetRouter()
	r.Run(utils.HttpPort)
}
