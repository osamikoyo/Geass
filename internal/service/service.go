package service

import (
	"fmt"
	"net/http"

	"github.com/osamikoyo/geass/pkg/loger"
	"golang.org/x/net/html"
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
	if n.Type == html.ElementNode && n.Data == "li" {
		var name string
		var price string
		var imageURL string

		for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Data {
			case "h2":
			if c.FirstChild != nil && c.FirstChild.Type == html.TextNode {
				name = c.FirstChild.Data
			}
			case "span":
			for _, _ = range c.Attr {
				if c.FirstChild != nil && c.FirstChild.Type == html.TextNode {
						price = c.FirstChild.Data
				}
			}
			case "img":
			for _, a := range c.Attr {
				if a.Key == "src" {
					imageURL = a.Val
				}
			}
			}
		}
		fmt.Println("Product Name:", name)
		fmt.Println("Price:", price)
		fmt.Println("Image URL:", imageURL)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.traverse(c, url)
	}
}

func (s *Service) Start(u string) error {
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

	return nil
}

func (s *Service) DisplayContent() {
	for i, u := range s.Contents {
		fmt.Printf("Url: %s Content: %s\n", i, u)
	}
}