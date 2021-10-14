package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var f func(string) string
	shaPtr := flag.String("sha", "256", "SHA type: 256, 384, or 512.")
	flag.Parse()
	sha := *shaPtr
	if sha == "512" {
		fmt.Println("Use SHA512")
		f = s512
	} else if sha == "384" {
		fmt.Println("Use SHA384")
		f = s384
	} else {
		fmt.Println("Use SHA256")
		f = s256
	}

	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			s := f(sc.Text())
			fmt.Println(s)
		}
	}
}

func s256(a string) string {
	bytes := sha256.Sum256([]byte((a)))
	return fmt.Sprintf("%x", bytes)
}

func s384(a string) string {
	bytes := sha512.Sum384([]byte((a)))
	return fmt.Sprintf("%x", bytes)
}

func s512(a string) string {
	bytes := sha512.Sum512([]byte((a)))
	return fmt.Sprintf("%x", bytes)
}
