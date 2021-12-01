package api

import (
	v1 "github.com/czasg/auth/src/interface/http/api/v1"
	v2 "github.com/czasg/auth/src/interface/http/api/v2"
	"github.com/kataras/iris/v12"
)

func RegisterRoute(party iris.Party) {
	v1.RegisterRoute(party.Party("/v1"))
	v2.RegisterRoute(party.Party("/v2"))
}
