// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 365.

//!+

// Bzipper reads input, bzip2-compresses it, and writes it out.
package main

import (
	"io"
	"log"
	"os"

	"go-hello/ch13/ex03/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	f, _ := os.Open("/usr/share/dict/words")
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}

//!-
