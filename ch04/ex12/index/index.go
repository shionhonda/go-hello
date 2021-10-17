package index

import (
	"encoding/json"
	"go-hello/ch04/ex12/mytypes"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func Index(comics *[]mytypes.Comic) *map[string][]mytypes.Comic {
	m := make(map[string][]mytypes.Comic)
	for _, c := range *comics {
		ws := split(strings.ToLower(c.Transcript))
		for _, w := range ws {
			cs, ok := m[w]
			if ok {
				m[w] = append(cs, c)
			} else {
				m[w] = []mytypes.Comic{c}
			}
		}
	}
	return &m
}

func LoadJSON(filepath string) *mytypes.Comic {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var comic mytypes.Comic
	json.Unmarshal(data, &comic)
	return &comic
}

func split(s string) []string {
	r := regexp.MustCompile("[^\\s,.!?]+")
	return r.FindAllString(s, -1)
}
