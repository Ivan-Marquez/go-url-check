package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type logger struct {
	data string
	url  string
	err  error
}

func fetchURL(duration time.Duration, urls []string) error {
	start := time.Now()
	end := start.Add(duration)

	ch := make(chan logger)

	for _, url := range urls {
		url = strings.TrimSpace(url)
		go fetch(url, ch)
	}

	// receive from channel ch
	for t := range ch {
		fmt.Println(t.data)
		switch {
		// close ch if end time reached
		case end.Before(time.Now()):
			close(ch)
		// close ch and return any network error
		case t.err != nil:
			close(ch)
			return t.err
		// if there's still time left
		// before specified end time
		case time.Now().Before(end):
			go func(url string) {
				if time.Now().Add(2 * time.Second).Before(end) {
					time.Sleep(2 * time.Second)
					fetch(url, ch)
				} else {
					close(ch)
				}
			}(t.url)
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	return nil
}

func fetch(url string, ch chan<- logger) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- logger{
			data: fmt.Sprint(err),
			err:  err,
		}
		return
	}

	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- logger{
			data: fmt.Sprintf("while reading %s: %v", url, err),
		}
		return
	}

	secs := time.Since(start).Seconds()
	ch <- logger{
		data: fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url),
		url:  url,
	}
}
