package main

import (
	"bufio"
	"fmt"
	"go-hello/ch04/ex12/download"
	"go-hello/ch04/ex12/index"
	"go-hello/ch04/ex12/mytypes"
	"log"
	"os"
)

const numComics int = 100

func main() {
	// Download data
	os.Mkdir("ex12/data", 0777)
	for i := 1; i <= numComics; i++ {
		filepath := fmt.Sprintf("ex12/data/%04d.json", i)
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			out, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			download.Download(out, i)
		} else {
			fmt.Printf("file exists: %s\n", filepath)
		}
	}

	// Index
	comics := []mytypes.Comic{}
	for i := 1; i <= numComics; i++ {
		filepath := fmt.Sprintf("ex12/data/%04d.json", i)
		comic := index.LoadJSON(filepath)
		comics = append(comics, *comic)
	}
	m := index.Index(&comics)

	// Search
	fmt.Println("Input query words:")
	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			search(sc.Text(), m)
		}
	}

}

func search(word string, m *map[string][]mytypes.Comic) {
	comics, ok := (*m)[word]
	if ok {
		for _, comic := range comics {
			url := fmt.Sprintf("https://xkcd.com/%d/", comic.Num)
			fmt.Printf("\n---\n%s\n%s\n%s\n", comic.Title, url, comic.Transcript)
		}
	} else {
		fmt.Println("Got 0 results.")
	}
}
