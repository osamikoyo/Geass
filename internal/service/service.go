package service

import (
	"io"
	"sync"

	"github.com/osamikoyo/geass/pkg/loger"
)

type Service struct {
	Logger loger.Logger
	URLS []string
	Contents map[string]string
}


func (s *Service) Start(u string, w io.Writer) error {
	var wg sync.WaitGroup

	wg.Add(1)
	go ParsePage(u, 1, &wg, w)

	wg.Wait()
	return nil
}

