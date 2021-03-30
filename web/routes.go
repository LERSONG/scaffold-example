package web

import "net/http"

func (srv *server) routes() http.Handler {
	srv.router.GET("scaffold-example/version", srv.Version)

	//Declare web routing table at here.
	return srv.router
}
