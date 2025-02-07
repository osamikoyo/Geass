package service

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func extractText(n *html.Node) string {
	var fulltext string

	// Если текущий узел является текстовым, добавляем его содержимое
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fulltext += text + " "
		}
	}

	// Рекурсивно обходим дочерние узлы
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fulltext += extractText(c)
	}

	return fulltext
}

// TextContentParse загружает HTML-страницу и извлекает текст
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