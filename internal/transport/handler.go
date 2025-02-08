package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/osamikoyo/geass/internal/service"
	"github.com/osamikoyo/geass/pkg/loger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	logger loger.Logger
	service *service.Service
}

type TextResponse struct{
	Url string `json:"url"`
	Text string `json:"text"`
}

func (h Handler) RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/get/content", h.ErrorRoute(h.GetContentHandler))
	mux.HandleFunc("/get/urls", h.ErrorRoute(h.GetUrlsHandler))
	mux.HandleFunc("/ping",  h.ErrorRoute(h.PingHandler))
	mux.HandleFunc("/get/text", h.ErrorRoute(h.GetPageTextContentHandler))
	mux.Handle("/metrics", promhttp.Handler())
}

func New(logsdir string) Handler {
	return Handler{
		service: &service.Service{
			Logger: loger.New(logsdir),
			URLS: make([]string, 1),
			Contents: make(map[string]string),
		},
	}
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h *Handler) ErrorRoute(handler handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(requestDuration)
		defer timer.ObserveDuration()
		if err := handler(w, r);err != nil{
			h.logger.Error().Str("URL", r.URL.Path).Str("METHOD", r.Method).Err(err)
		}
		requestsCount.Inc()
	}
}

func (h *Handler) GetPageTextContentHandler(w http.ResponseWriter, r *http.Request) error {
	timer := prometheus.NewTimer(requestDuration)
	defer timer.ObserveDuration()
	url := r.URL.Query().Get("url")

	text, err := h.service.TextContentParse(url)
	if err != nil{
		return err
	}

	resp := TextResponse{
		Text: text,
		Url: url,
	}

	w.Header().Set("Content-type", "application/json")
	body, err := json.Marshal(resp)
	if err != nil{
		return err
	}

	requestsCount.Inc()
	_, err = w.Write(body)
	return err
}

func (h *Handler) GetContentHandler(w http.ResponseWriter, r *http.Request) error {
	url := r.URL.Query().Get("url")

	pageinfo, err := h.service.ContentParsePage(url)
	if err != nil{
		h.logger.Info().Str("URL", url).Msg("Cant Parse!")
		http.Error(w, "cant parse: " + err.Error(), http.StatusInternalServerError)
		return err
	}

	respbody, err := json.Marshal(pageinfo)
	if err != nil{
		h.logger.Info().Str("URL", url).Msg("Cant marshal!: " + err.Error())
		http.Error(w, "cant marshal: " + err.Error(), http.StatusInternalServerError)
		return err
	}

	fmt.Fprint(w, string(respbody))
	return nil
}

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "PING!!:3")
	return nil
}

func (h *Handler) GetUrlsHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return errors.New("invalud request method")
	}

	pageURL := r.URL.Query().Get("url")
	if pageURL == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return errors.New("url = nil")
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Отправляем ответ в формате текста
	w.Header().Set("Content-Type", "text/plain")
	go service.ParsePage(pageURL, 0, &wg, w)

	wg.Wait()
	return nil	
}