package Init

import (
	"learn-gin/Packages/Core/Route"
	"learn-gin/Routes"
)

func Bootstrap() {
	//Db.Connect()
	Route.InitRoute()
	Routes.RegisterAPIRoutes(Route.InitRoute())
}
