package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/net/html"
)

type visitedURLs struct {
	mu   sync.Mutex
	urls map[string]bool
}

func (v *visitedURLs) add(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.urls[url] = true
}

func (v *visitedURLs) isVisited(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.urls[url]
}

var visited = visitedURLs{urls: make(map[string]bool)}

func extractLinks(n *html.Node, baseURL string, depth int, wg *sync.WaitGroup, w io.Writer) {
    defer wg.Done()

    if n.Type == html.ElementNode {
        switch n.Data {
        case "a":
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    link := attr.Val
                    absoluteURL := resolveURL(baseURL, link)

                    if !visited.isVisited(absoluteURL) {
                        visited.add(absoluteURL)
                        fmt.Println("Ссылка:", absoluteURL)
						w.Write([]byte(absoluteURL))
                        if depth < 2 {
                            wg.Add(1)
                            go parsePage(absoluteURL, depth+1, wg, w)
                        }
                    }
                    break
                }
            }
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        wg.Add(1)
        go extractLinks(c, baseURL, depth, wg, w)
    }
}

func parsePage(url string, depth int, wg *sync.WaitGroup, w io.Writer) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	wg.Add(1)
	extractLinks(doc, url, depth, wg, w)
}

func resolveURL(baseURL, link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	return base.ResolveReference(u).String()
}
