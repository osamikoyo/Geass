package server

import (
	"github.com/osamikoyo/geass/internal/transport"
	"github.com/osamikoyo/geass/pkg/loger"
	"net/http"
)

type Server struct {
	HttpServer *http.Server
	Logger loger.Logger
	Handler *transport.Handler
	mux *http.ServeMux
}

func New() Server {
	handler := transport.New([]string{"https://ru.wikipedia.org/wiki/Motion_blur"})
	return Server{
		Handler: &handler,
		Logger: loger.New(),
		HttpServer: &http.Server{
			Addr: "localhost:8080",
		},
		mux: http.NewServeMux(),
	}
}

func (s *Server) Run() {
	
}