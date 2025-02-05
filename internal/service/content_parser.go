package service

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// Структуры данных
type Image struct {
	Src string
	Alt string
}

type Link struct {
	Text string
	Href string
}

type Content struct {
	FullText string
	Images   []Image
}

type Technical struct {
	Code        uint32
	ContentType string
}

type Metadata struct {
	Lang   string
	Robots string
}

type PageInfo struct {
	Url               string
	Title            string
	MetadataDescription string
	Content          Content
	CountKeyWord     uint64
	Links            []Link
	Technical        Technical
	Metadata         Metadata
}

func (s *Service)ContentParsePage(url string) (*PageInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке страницы: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при парсинге HTML: %v", err)
	}

	pageInfo := &PageInfo{
		Url: url,
		Technical: Technical{
			Code:        uint32(resp.StatusCode),
			ContentType: resp.Header.Get("Content-Type"),
		},
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil {
					pageInfo.Title = n.FirstChild.Data
				}
			case "meta":
				if getAttr(n, "name") == "description" {
					pageInfo.MetadataDescription = getAttr(n, "content")
				}
				if getAttr(n, "name") == "robots" {
					pageInfo.Metadata.Robots = getAttr(n, "content")
				}
			case "img":
				src := getAttr(n, "src")
				alt := getAttr(n, "alt")
				if src != "" {
					pageInfo.Content.Images = append(pageInfo.Content.Images, Image{Src: src, Alt: alt})
				}
			case "a":
				href := getAttr(n, "href")
				text := getText(n)
				if href != "" {
					pageInfo.Links = append(pageInfo.Links, Link{Text: text, Href: href})
				}
			case "html":
				pageInfo.Metadata.Lang = getAttr(n, "lang")
			}
		} else if n.Type == html.TextNode {
			pageInfo.Content.FullText += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	pageInfo.CountKeyWord = uint64(len(strings.Fields(pageInfo.Content.FullText)))

	return pageInfo, nil
}

func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getText(n *html.Node) string {
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
		}
	}
	return strings.TrimSpace(text)
}