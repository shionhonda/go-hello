package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	filename, len, err := fetch(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("downloaded %s: %d bytes\n", filename, len)
}

// fetch downloads the given URL and returns its filename and bytes of content.
func fetch(url string) (filename string, bytes int64, retErr error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	// check if the file is successfylly closed and return it if no other error happened
	defer func() {
		if closeErr := f.Close(); retErr == nil {
			retErr = closeErr
		}
	}()
	n, err := io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, err
	}

	return local, n, err
}
