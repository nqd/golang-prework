package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	request := flag.Int64("n", 1, "Number of request to perform")
	concurrency := flag.Int64("c", 1, "Number of multiple requests to make at a time")

	flag.Parse()

	if flag.NArg() == 0 || *request <= 0 || *concurrency <= 0 || *request < *concurrency {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	url := flag.Arg(0)
	per(url)

}

func per(url string) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	read, _ := io.Copy(ioutil.Discard, res.Body)

	fmt.Println(read, res.StatusCode)

}
