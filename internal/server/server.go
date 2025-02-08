package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/osamikoyo/geass/internal/config"
	"github.com/osamikoyo/geass/internal/transport"
	"github.com/osamikoyo/geass/pkg/loger"
)

type Server struct {
	HttpServer *http.Server
	Logger loger.Logger
	Handler *transport.Handler
	mux *http.ServeMux
	config *config.Config
}
func New() Server {
	config, _ := config.Get("config.yml")

	handler := transport.New(config.LogsDir)
	return Server{
		Handler: &handler,
		Logger: loger.New(config.LogsDir),
		HttpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		},
		mux: http.NewServeMux(),
		config: config,
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


	transport.InitMetrix()

	s.Handler.RegisterRouter(s.mux)

	s.HttpServer.Handler = s.mux

	fmt.Print(`
  ________                             
 /  _____/  ____ _____    ______ ______
/   \  ____/ __ \\__  \  /  ___//  ___/
\    \_\  \  ___/ / __ \_\___ \ \___ \ 
 \______  /\___  >____  /____  >____  >
        \/     \/     \/     \/     \ 


`)
	s.Logger.Info().Str("Server Addr", s.HttpServer.Addr).Msg("SERVER STARTING!")
	
	if err := s.HttpServer.ListenAndServe();err != nil{
		return err
	}

	s.Logger.Info().Msg("ALWAYS WORK CORRECT")
	return nil
}