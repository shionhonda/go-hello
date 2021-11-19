package main

import (
	"go-hello/ch07/ex08/sorting"
	"html/template"
	"log"
	"os"
)

type Result struct {
	TotalCount int
	Items      []*sorting.Track
}

var trackList = template.Must(template.New("trackList").Parse(`
<h1>{{.TotalCount}} Tracks</h1>
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>Artist</th>
  <th>Album</th>
  <th>Year</th>
  <th>Length</th>
</tr>
{{range .Items}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	res := Result{
		TotalCount: len(sorting.Tracks),
		Items:      sorting.Tracks,
	}

	if err := trackList.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}
