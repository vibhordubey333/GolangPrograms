package main

import (
	"fmt"
	"net/http"
	"time"
)

var count int64

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler Says Hello")
	time.Sleep(500 * time.Millisecond)
	w.WriteHeader(200)
	r.Body.Close()
	count++
}

func main() {
	http.HandleFunc("/counter", handle)
	fmt.Println("Starting Server:")
	if err := http.ListenAndServe(":7777", nil); err != nil {
		panic(err)
	}
}

/*
========Apache AB==========
Output:
ab -c 500 -n 500 http://localhost:7777/counter
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Finished 500 requests


Server Software:
Server Hostname:        localhost
Server Port:            7777

Document Path:          /counter
Document Length:        0 bytes

Concurrency Level:      500
Time taken for tests:   1.220 seconds
Complete requests:      500
Failed requests:        0
Total transferred:      37500 bytes
HTML transferred:       0 bytes
Requests per second:    409.71 [#/sec] (mean)
Time per request:       1220.388 [ms] (mean)
Time per request:       2.441 [ms] (mean, across all concurrent requests)
Transfer rate:          30.01 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       4
Processing:   500  513  10.9    513     710
Waiting:      500  512   6.5    512     528
Total:        501  513  10.9    513     710

Percentage of the requests served within a certain time (ms)
  50%    513
  66%    516
  75%    518
  80%    519
  90%    522
  95%    524
  98%    525
  99%    527
 100%    710 (longest request)

*/
