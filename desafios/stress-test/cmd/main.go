package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	url         string
	concurrency int
	requests    int
)

func main() {
	flag.StringVar(&url, "url", "https://google.com", "url to be stressed")
	flag.IntVar(&concurrency, "concurrency", 20, "number of concurrent requests")
	flag.IntVar(&requests, "requests", 3500, "number of requests")
	flag.Parse()

	logrus.Info("url ", url, " requests ", requests, " concurrency ", concurrency)

	wg := sync.WaitGroup{}
	start := time.Now()
	flowCh := make(chan struct{}, concurrency)
	errorCH := make(chan int)
	statusCH := make(chan int)

	result := make(map[string]int)
	resultErr := make(map[string]int)

	go func() {
		for {
			select {
			case <-errorCH:
				resultErr["errors"] += 1
			case s := <-statusCH:
				result[strconv.Itoa(s)] += 1
			}
		}
	}()

	for i := 0; i < requests; i++ {
		flowCh <- struct{}{}
		wg.Add(1)
		go makeRequests(&wg, flowCh, errorCH, statusCH, url)
	}
	wg.Wait()
	logrus.Info("Execution time: ", time.Since(start))
	logrus.Info("Total requests: ", requests)
	logrus.Info("Total errors requests: ", resultErr["errors"])
	logrus.Info("Total success requests: ", requests-resultErr["errors"])
	for k, v := range result {
		logrus.Info("Status: ", k, ": total: ", v)
	}
}

func makeRequests(wg *sync.WaitGroup, flowCH <-chan struct{}, errorCH chan<- int, statusCH chan<- int, url string) {
	defer wg.Done()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		errorCH <- 1
		<-flowCH
		return
	}

	statusCH <- response.StatusCode
	<-flowCH
}
