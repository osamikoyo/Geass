package app

import (
	"github.com/osamikoyo/geass/internal/service"
	"github.com/osamikoyo/geass/pkg/loger"
)

func Init() App {
	return App{
		logger: loger.New(),
		Service: service.Service{
			Logger: loger.New(),
			URLS: []string{"https://ru.wikipedia.org/wiki/Motion_blur"},
		},
	}
}

type App struct {
	Service service.Service
	logger loger.Logger
}