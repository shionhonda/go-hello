Popcount with `x&(x-1)` trick is a little faster!

```console
# Time for 1,000,000 loops
$ go run ch02/ex05/main.go 123456
Base: 0.004690s
Trick: 0.004282s
```