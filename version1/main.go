package main

import (
	"version1/model"
	"version1/router"
	"version1/utils"
)

func main() {
	utils.Init()
	model.InitDB()
	router.SetRouter()
}
