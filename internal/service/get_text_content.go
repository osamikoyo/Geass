package service

import (
	"fmt"

	"golang.org/x/net/html"
)

func (s *Service) GetTextContent(n *html.Node) (string, error) {
	var text string

	if n.Type == html.ElementNode {
		if n.Data == "title" {
			if n.FirstChild != nil{
				text = fmt.Sprintf("%s%s\n\n", text, n.FirstChild.Data)
			}
		}
		if n.Data == "h1" || n.Data == "h2" || n.Data == "h3"||
		n.Data == "h4"  || n.Data == "h5" || n.Data == "h6" ||
		n.Data == "a" || n.Data == "span" || n.Data == "p"{
			if n.FirstChild != nil {
				text = fmt.Sprintf("%s%s\n", text, n.FirstChild.Data)
			}
		}

	}

	return text,nil
}