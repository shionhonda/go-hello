package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	xmlText := fetch(url)
	fmt.Println("Names?")
	names := askStrs()
	fmt.Println("Attributes?")
	attrs := askMap()

	fmt.Println("**search condition**")
	for i, name := range names {
		fmt.Printf("name %d: %s\n", i, name)
	}
	for k, v := range attrs {
		fmt.Printf("map %s: %v\n", k, v)
	}

	dec := xml.NewDecoder(strings.NewReader(xmlText))
	var stack []string // stack of element names
	var astack map[string]string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			astack = make(map[string]string)
			for _, a := range tok.Attr {
				astack[a.Name.Local] = a.Value
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, names) && matchAll(astack, attrs) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// matchAll reports whether x contains the keys and values of y
func matchAll(x, y map[string]string) bool {
	for k, v := range y {
		if x[k] != v {
			return false
		}
	}
	return true
}

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s", b)
}

func askStrs() []string {
	var texts []string
	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			texts = append(texts, sc.Text())
		} else {
			break
		}
	}
	return texts
}

func askMap() map[string]string {
	attrs := make(map[string]string)
	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			strs := strings.Split(sc.Text(), ":")
			if len(strs) == 2 {
				attrs[strs[0]] = strs[1]
			} else {
				fmt.Printf("skip: %v\n", strs)
				continue
			}
		} else {
			break
		}
	}
	return attrs
}
