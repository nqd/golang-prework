# CoderSchool Golang Course Prework Submission Template

[![Build Status](https://travis-ci.org/nqd/golang-prework.svg?branch=master)](https://travis-ci.org/nqd/golang-prework)

1. **Submitted by: Nguyen Quoc Dinh**
2. **Time spent:**

## Set of User Stories

### Required

* [x] Command-line argument parsing
* [x] Input params
  * [x] Requests - Number of requests to perform
  * [x] Concurrency - Number of multiple requests to make at a time
  * [x] URL - The URL for testing.
* [x] Prints use information if wrong arguments provided
* [x] Implements  HTTP load and summarize it
* [x] Concurrency must be implemented with goroutine.

### Bonus

* [x] Extend input params with:
  * [x] Timeout - Seconds to max. wait for each response
  * [x] Timelimit - Maximum number of seconds to spend for benchmarking
* [ ] Prints key metrics of summary, such:
  * [ ] Server Hostname
  * [ ] Server Port
  * [ ] Document Path
  * [ ] Document Length
  * [ ] Concurrency Level
  * [ ] Time taken for tests
  * [ ] Complete requests
  * [ ] Failed requests
  * [ ] Total transferred
  * [ ] Requests per second
  * [ ] Time per request
  * [ ] Time per request
  * [ ] Transfer rate

### Test

`go run main.go -n 30 -c 2 -s 10 -t 5 https://bbc.co.uk`