package main

import (
	"bytes"
	"fmt"
	"go-hello/ch12/ex07/sexpr"
	"log"
	"reflect"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// encode
	data, err := sexpr.Marshal(strangelove)
	if err != nil {
		panic(err)
	}

	// decode
	decoder := sexpr.NewDecoder(bytes.NewReader(data))
	var movie Movie
	err = decoder.Decode(&movie)
	if err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}
	if !reflect.DeepEqual(movie, strangelove) {
		log.Fatal("not equal")
	} else {
		fmt.Println("successfully reconstructed")
	}
}
