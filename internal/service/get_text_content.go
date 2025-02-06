package service

import (
	"fmt"

	"golang.org/x/net/html"
)

const enter = `

`

func (s *Service) GetTextContent(n *html.Node) (string, error) {
	var text string

	if n.Type == html.ElementNode {
		switch n.Data{
		case "title":
			if n.FirstChild != nil {
				text = fmt.Sprintf("%s%s\n\n", text, n.FirstChild.Data)
			}
		
		}
	}
}