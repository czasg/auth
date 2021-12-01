package auth

import (
	"github.com/czasg/auth/src/interface/http/api"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	api.RegisterRoute(app.Party("/api"))
	_ = app.Listen(":8080", iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           true,
		DisableBodyConsumptionOnUnmarshal: true,
		DisableStartupLog:                 true,
		RemoteAddrHeaders: map[string]bool{
			"X-Real-Ip":        true,
			"X-Forwarded-For":  true,
			"CF-Connecting-IP": true,
		},
	}))
}
