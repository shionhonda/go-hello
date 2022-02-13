`MaxResults` is capped at 1000.

```sh
$ go run main.go
$ curl "http://localhost:12345/search?l=golang&l=programming&max=2000"
Search: {Labels:[golang programming] MaxResults:1000 Exact:false}
```