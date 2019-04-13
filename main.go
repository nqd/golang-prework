package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type perInfo struct {
	err      error
	status   int
	bytes    int64
	duration time.Duration
}

func main() {
	request := flag.Int64("n", 1, "Number of request to perform")
	concurrency := flag.Int64("c", 1, "Number of multiple requests to make at a time")
	s := flag.Int64("s", 30, "Maximum number of seconds to wait before the socket times out")
	t := flag.Int64("t", 50000, "Maximum number of seconds to spend for benchmarking")

	flag.Parse()

	if flag.NArg() == 0 || *request <= 0 || *concurrency <= 0 ||
		*request < *concurrency ||
		*s < 0 ||
		*t < 0 {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	url := flag.Arg(0)
	timeout := time.Second * time.Duration(*s)
	timelimit := time.Second * time.Duration(*t)

	result := make(chan perInfo)
	for i := int64(0); i < *concurrency; i++ {
		go per(url, timeout, result)
	}

	for i := int64(0); i < *request; i++ {
		if i+*concurrency < *request {
			go per(url, timeout, result)
		}

		res := <-result
		log.Println(res)
	}
}

func per(url string, timeout time.Duration, result chan perInfo) {
	start := time.Now()

	client := &http.Client{
		Timeout: timeout,
	}
	res, err := client.Get(url)

	// maybe from the timeout
	if err != nil {
		result <- perInfo{
			err: err,
		}
		return
	}
	defer res.Body.Close()

	bytes, _ := io.Copy(ioutil.Discard, res.Body)

	duration := time.Now().Sub(start)

	result <- perInfo{
		bytes:    bytes,
		status:   res.StatusCode,
		duration: duration,
	}
}
