package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/osamikoyo/geass/internal/transport"
	"github.com/osamikoyo/geass/pkg/loger"
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

func (s *Server) Shutdown(ctx context.Context) {
	s.Logger.Info().Msg("Server shutdown!")
	s.HttpServer.Shutdown(ctx)
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.Shutdown(ctx)
	}()


	s.Handler.RegisterRouter(s.mux)

	s.HttpServer.Handler = s.mux
	s.Logger.Info().Str("Server Addr", s.HttpServer.Addr).Msg("SERVER STARTING!")

	fmt.Print(`
  ________                             
 /  _____/  ____ _____    ______ ______
/   \  ____/ __ \\__  \  /  ___//  ___/
\    \_\  \  ___/ / __ \_\___ \ \___ \ 
 \______  /\___  >____  /____  >____  >
        \/     \/     \/     \/     \/ `)
	
	if err := s.HttpServer.ListenAndServe();err != nil{
		return err
	}

	s.Logger.Info().Msg("ALWAYS WORK CORRECT")
	return nil
}