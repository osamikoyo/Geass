package transport

import (
	"github.com/osamikoyo/geass/internal/service"
	"github.com/osamikoyo/geass/pkg/loger"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func New(urls []string) Handler {
	return Handler{
		service: &service.Service{
			Logger: loger.New(),
			URLS: urls,
			Contents: make(map[string]string),
		},
	}
}

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) error {
	
}