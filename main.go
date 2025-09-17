package main

import (
	"learn-gin/Packages/Core/Init"
	"learn-gin/Packages/Core/Route"
)

func main() {
	Init.Bootstrap()
	Route.InitRoute().Run(":7002")

}
