package Server

import (
	"net/http"
	"github.com/dapsframework/daps/Foundation"
)

type Server struct {
	Pattern		string
	Uri 		string
	Handler 	http.Handler
	Application 	*Foundation.Application
}

func (server *Server) Serve() {
	http.Handle(server.Pattern, server.Handler)
	http.ListenAndServe(server.Uri, nil)
}
