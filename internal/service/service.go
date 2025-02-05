package service

import (
	"fmt"
	"sync"

	"github.com/osamikoyo/geass/pkg/loger"
)

type Service struct {
	Logger loger.Logger
	URLS []string
	Contents map[string]string
}


func (s *Service) Start(u string) error {
	var wg sync.WaitGroup

	wg.Add(1)
	go parsePage(u, 1, &wg)

	wg.Wait()
	return nil
}

func (s *Service) DisplayContent() {
	for i, u := range s.Contents {
		fmt.Printf("Url: %s Content: %s\n", i, u)
	}
}