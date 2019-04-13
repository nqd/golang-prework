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
	status   int
	bytes    int64
	duration time.Duration
}

func main() {
	request := flag.Int64("n", 1, "Number of request to perform")
	concurrency := flag.Int64("c", 1, "Number of multiple requests to make at a time")
	timeout := flag.Int64("t", 30, "Maximum number of seconds to wait before the socket times out")

	flag.Parse()

	if flag.NArg() == 0 || *request <= 0 || *concurrency <= 0 ||
		*request < *concurrency ||
		*timeout < 0 {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	url := flag.Arg(0)

	result := make(chan perInfo)
	for i := int64(0); i < *concurrency; i++ {
		go per(url, result)
	}

	for i := int64(0); i < *request; i++ {
		if i+*concurrency < *request {
			go per(url, result)
		}

		res := <-result
		log.Println(res)
	}
}

func per(url string, result chan perInfo) {
	start := time.Now()
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	bytes, _ := io.Copy(ioutil.Discard, res.Body)

	duration := time.Now().Sub(start)

	result <- perInfo{
		bytes:    bytes,
		status:   res.StatusCode,
		duration: duration,
	}
}
