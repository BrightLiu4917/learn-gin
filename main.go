// main.go
package main

import (
	"learn-gin/Packages/Core/Init"
	"learn-gin/Packages/Core/Route"
)

func main() {
	// 初始化系统
	Init.Bootstrap()

	// 启动服务
	Route.GetRouter().Run(":7002")
}
