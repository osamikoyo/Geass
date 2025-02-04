package transport

import (
	"errors"
	"net/http"

	"github.com/osamikoyo/geass/internal/service"
	"github.com/osamikoyo/geass/pkg/loger"
)

type Handler struct {
	logger loger.Logger
	service *service.Service
}

func (h *Handler) GET(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			h.logger.Error().Str("URL", r.URL.Path).Err(errors.New(
				"method not get",
			))
		}
	}
}

func (h *Handler) RegisterRouter(mux *http.ServeMux) {
	mux.Handle("/getcontent", h.GET(h.ErrorRoute(h.MainHandler)))
	mux.Handle("/ping",  h.GET(h.ErrorRoute(h.PingHandler)))
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

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h *Handler) ErrorRoute(handler handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r);err != nil{
			h.logger.Error().Str("URL", r.URL.Path).Str("METHOD", r.Method).Err(err)
		}
	}
}

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("PONG!!!:333"))
	return err
}

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) error {
	url := r.URL.Query().Get("url")

	h.service.AddUrl(url)
	h.service.Start()
	h.service.DisplayContent()
	return nil	
}