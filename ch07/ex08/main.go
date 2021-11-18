package main

import (
	"fmt"
	"go-hello/ch07/ex08/sorting"
	"sort"
)

func printTrack(t *sorting.Track) {
	fmt.Printf("%s | %s | %s | %d | %s\n", t.Title, t.Artist, t.Album, t.Year, t.Length)
}

func main() {
	mc := sorting.NewMultiColumn(sorting.Tracks, [3]string{"Title", "Artist", "Album"})
	fmt.Println("Before sorting:")
	mc.PrintTracks()

	sort.Sort(mc)
	fmt.Println("\nAfter sorting:")
	mc.PrintTracks()
}
