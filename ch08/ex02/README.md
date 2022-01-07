# FTP Server
## Server
```sh
go run server/server.go
```

## Client
`inetutils`'s `ftp` command is not supported somehow.

```sh
$ ftp localhost
ftp: connect to address ::1: Connection refused
ftp: Trying 127.0.0.1 ...
ftp: connect to address 127.0.0.1: Connection refused
ftp: no response from host
ftp>
[4]+  Stopped                 ftp localhost
```

So, use the following client instead:

```sh
$ go run client/client.go
ls
README.md
client
server
cd server
ls
server.go
get server.g
open server.g: no such file or directory
get server.go
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")
		switch cmds[0] {
		case "ls":
			dir := "."
			if len(cmds) > 1 {
				dir = cmds[1]
			}
			files, err := os.ReadDir(dir)
			if err != nil {
				mustCopy(c, strings.NewReader(string(err.Error())+"\n"))
				continue
			}
			filenames := make([]string, len(files))
			for i, file := range files {
				filenames[i] = file.Name()
			}
			mustCopy(c, strings.NewReader(strings.Join(filenames, "\n")+"\n"))
		case "cd":
			if len(cmds) < 2 {
				mustCopy(c, strings.NewReader("cd: no destination\n"))
				continue
			}
			if err := os.Chdir(cmds[1]); err != nil {
				mustCopy(c, strings.NewReader(string(err.Error())+"\n"))
			}
		case "get":
			if len(cmds) < 2 {
				mustCopy(c, strings.NewReader("get: no file specified\n"))
				continue
			}
			file, err := os.Open(cmds[1])
			if err != nil {
				mustCopy(c, strings.NewReader(string(err.Error())+"\n"))
				continue
			}
			defer file.Close()
			mustCopy(c, file)
		case "close":
			return
		default:
			help := "ls: list content\ncd: change directory\nget: get file content\nclose: close connection\n"
			mustCopy(c, strings.NewReader(help))
		}
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
close
```