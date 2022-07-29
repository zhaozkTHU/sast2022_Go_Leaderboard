package main

import (
	"leadboard/config"
	"leadboard/model"
	"leadboard/route"
)

func main() {
	r := route.InitRoute()
	model.BuildConnection(config.Parse())
	r.Run(":8080")
}
