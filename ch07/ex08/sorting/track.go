package sorting

import "time"

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var Tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"We Got The Beat", "Go-Go's", "Beauty and the Beat", 1981, length("2m26s")},
	{"Go", "Asia", "Astra", 1985, length("3m40s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
