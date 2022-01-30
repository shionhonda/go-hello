```sh
$ go run main.go 
Display c (main.Cycle):
c.Value = 42
(*c.Tail).Value = 42
(*(*c.Tail).Tail).Value = 42
(*(*(*c.Tail).Tail).Tail).Value = 42
(*(*(*(*c.Tail).Tail).Tail).Tail).Value = 42
```