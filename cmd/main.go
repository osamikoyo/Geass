package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/osamikoyo/geass/internal/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := server.New()
	if err := s.Run(ctx);err != nil{
		s.Logger.Error().Err(err)
	}
}