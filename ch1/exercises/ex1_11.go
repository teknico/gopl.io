// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.11, p. 19: try fetchall with longer argument lists, such as
// samples from the top million web sites available at alexa.com. How does
// the program behave if a web site just doesn’t respond? (Section 8.9
// describes mechanisms for coping in such cases.)
// Based on fetchall.
/*
cat ex1_11-urls.txt | xargs go run ex1_11.go
Get http://1e100.net: dial tcp: lookup 1e100.net on 127.0.1.1:53: no such host
Get http://bp.blogspot.com: dial tcp: lookup bp.blogspot.com on 127.0.1.1:53: no such host
Get http://about.com: unexpected EOF
2.07s    38250  http://ameblo.jp
...
5.42s    83265  http://4shared.com
Get http://orkut.com.br: dial tcp: lookup orkut.com.br on 127.0.1.1:53: server misbehaving
5.63s   186280  http://alibaba.com
...
9.83s   175297  http://livejasmin.com
Get http://orkut.co.in: dial tcp: lookup orkut.co.in on 127.0.1.1:53: read udp 127.0.0.1:52507->127.0.1.1:53: i/o timeout
Get http://walmart.com: dial tcp: lookup walmart.com on 127.0.1.1:53: read udp 127.0.0.1:59160->127.0.1.1:53: i/o timeout
Get http://yandex.ru: dial tcp [2a02:6b8:a::a]:80: connect: network is unreachable
Get http://xhamster.com: dial tcp [2a02:b49:4:8::1]:80: connect: network is unreachable
10.05s   426727  http://sohu.com
...
14.64s    14078  http://cnzz.com
Get http://rapidshare.com: dial tcp 10.11.12.13:80: i/o timeout
Get http://yieldmanager.com: dial tcp 208.67.66.24:80: i/o timeout
48.79s    76366  http://kaixin001.com
while reading http://tudou.com: read tcp 192.168.1.155:47486->123.126.98.148:80: read: connection reset by peer
101.53s elapsed
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
