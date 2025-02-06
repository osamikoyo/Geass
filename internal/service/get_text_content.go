package service

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func extractText(n *html.Node) string {
	var fulltext string

	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fulltext = fmt.Sprintf("%s%s", fulltext, text)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c)
	}

	return fulltext
}

func (s *Service) TextContentParse(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	return extractText(doc), nil
}