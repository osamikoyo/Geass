package service

import (
	"fmt"
	"github.com/osamikoyo/geass/pkg/loger"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
)

type Service struct {
	Logger loger.Logger
	URLS []string
	Contents map[string]string
}

func (s *Service) AddUrl(url string) {
	s.URLS = append(s.URLS, url)
}

func (s *Service) traverse(n *html.Node, url string) {
	if n.Type == html.ElementNode {
		s.Contents[url] = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.traverse(c, url)
	}
}

func (s *Service) Start() error {
	for _, u := range s.URLS {

		s.Logger.Info().Str("URL", u).Msg("Started request")

		resp, err := http.Get(u)
		if err != nil{
			s.Logger.Error().Str("URL", u).Msg("Cant do a request: " + err.Error())
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil{
			s.Logger.Error().Str("URL", u).Msg("Cant parse to html: " + err.Error())
		}
		s.traverse(doc , u)
	}

	return nil
}