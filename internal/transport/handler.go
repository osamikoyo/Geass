package transport

import (
	"fmt"
	"net/http"

	"github.com/osamikoyo/geass/internal/service"
	"github.com/osamikoyo/geass/pkg/loger"
)

type Handler struct {
	logger loger.Logger
	service *service.Service
}

func (h Handler) RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/getcontent", h.ErrorRoute(h.MainHandler))
	mux.HandleFunc("/ping",  h.ErrorRoute(h.PingHandler))
}

func New() Handler {
	return Handler{
		service: &service.Service{
			Logger: loger.New(),
			URLS: make([]string, 1),
			Contents: make(map[string]string),
		},
	}
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h *Handler) ErrorRoute(handler handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r);err != nil{
			h.logger.Error().Str("URL", r.URL.Path).Str("METHOD", r.Method).Err(err)
		}
	}
}

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "PING!!:3")
	return nil
}

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) error {
	url := r.URL.Query().Get("url")

	h.service.Start(url)
	h.service.DisplayContent()
	return nil	
}