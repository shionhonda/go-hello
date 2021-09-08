package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := []string{"https://hippocampus-garden.com/best_papers_2020/"}
	for _, url := range urls {
		go fetch(url, ch, os.Args[1])
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, fname string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}

	dst, _ := os.Create(fname)
	defer dst.Close()

	nbytes, err := io.Copy(dst, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch: reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
