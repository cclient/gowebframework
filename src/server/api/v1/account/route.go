package account

import (
	"server/api/router"
	"server/api/v1/local"
)

type accountRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	r := &accountRouter{}
	r.initRoutes()
	return r
}

func (r *accountRouter) Routes() []router.Route {
	return r.routes
}

func (r *accountRouter) initRoutes() {
	r.routes = []router.Route{
		local.NewPostRoute("/api/v1/account/", r.InsertAccount),
		local.NewDeleteRoute("/api/v1/account/{id:.*}", r.RemoveAccountById),
		local.NewPutRoute("/api/v1/account/", r.UpdateAccountInfo),
		local.NewGetRoute("/api/v1/account/test", r.GetAccountTest),
		local.NewGetRoute("/api/v1/account/{id:.*}", r.GetAccountById),
		local.NewGetRoute("/api/v1/account/", r.GetAccountsPage),
	}
}
