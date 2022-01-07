package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[2:] {
		go time(arg)
	}
	time(os.Args[1])
}

func time(arg string) {
	strs := strings.Split(arg, "=")
	if len(strs) != 2 {
		log.Fatalf("invalid argument: %s", arg)
	}
	conn, err := net.Dial("tcp", strs[1])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn, strs[0])
}

func mustCopy(dst io.Writer, src io.Reader, name string) {
	var buf [16]byte
	for {
		n, _ := src.Read(buf[:])
		if n > 0 {
			str := fmt.Sprintf("%16s: %s", name, string(buf[:n]))
			dst.Write([]byte(str))
		}
	}
}
