package b

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, routes *mux.Router) error {

	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: routes,
	}
	return s.httpServer.ListenAndServe()
}
