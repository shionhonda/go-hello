package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Download(out io.Writer, num int) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)
}
