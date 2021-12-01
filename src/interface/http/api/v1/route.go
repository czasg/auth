package v1

import "github.com/kataras/iris/v12"

func RegisterRoute(party iris.Party) {
	party.Post("/user")
	party.Delete("/user")
}
