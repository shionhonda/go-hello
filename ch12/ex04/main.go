package main

import (
	"fmt"
	"go-hello/ch12/ex04/sexpr"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	// Bool            []bool
	// Float           float64
	// Complex         complex128
	// Interface       interface{}
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
		// Bool:    []bool{true, false},
		// Float:   -0.4,
		// Complex: 0.5 + 0.5i,
		// Interface: []int{1, 2, 3},
	}

	data, err := sexpr.MarshalIndent(strangelove)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
