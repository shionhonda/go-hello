package sorting

import "fmt"

type MultiColumn struct {
	t    []*Track
	keys [3]string
}

func NewMultiColumn(t []*Track, keys [3]string) *MultiColumn {
	return &MultiColumn{t, keys}
}

func (m *MultiColumn) Len() int {
	return len(m.t)
}

func (m *MultiColumn) Swap(i, j int) {
	m.t[i], m.t[j] = m.t[j], m.t[i]
}

func (m *MultiColumn) Less(i, j int) bool {
	for _, k := range m.keys {
		switch k {
		case "Title":
			if m.t[i].Title != m.t[j].Title {
				return m.t[i].Title < m.t[j].Title
			}
		case "Artist":
			if m.t[i].Artist != m.t[j].Artist {
				return m.t[i].Artist < m.t[j].Artist
			}
		case "Album":
			if m.t[i].Album != m.t[j].Album {
				return m.t[i].Album < m.t[j].Album
			}
		case "Year":
			if m.t[i].Year != m.t[j].Year {
				return m.t[i].Year < m.t[j].Year
			}
		case "Length":
			if m.t[i].Length != m.t[j].Length {
				return m.t[i].Length < m.t[j].Length
			}
		}
	}
	return false
}

func (m *MultiColumn) SetPrimary(key string) {
	for i, k := range m.keys {
		if k == key {
			copy(m.keys[i+1:], m.keys[:i])
			m.keys[0] = m.keys[i]
			return
		}
	}
	m.keys = [3]string{key, m.keys[0], m.keys[1]}

}

func (m *MultiColumn) PrintTracks() {
	for _, t := range m.t {
		fmt.Printf("%s | %s | %s | %d | %s\n", t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
}
