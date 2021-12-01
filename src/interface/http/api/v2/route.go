package v2

import "github.com/kataras/iris/v12"

func RegisterRoute(party iris.Party){
	party.Post("/login")
	party.Delete("/logout")
	party.Post("/status")
}
