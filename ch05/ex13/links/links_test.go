package links

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	var tests = []struct {
		url   string
		isErr bool
		list  []string
	}{
		{"https://www.gopl.io/", false, []string{
			"http://www.informit.com/store/go-programming-language-9780134190440",
			"http://www.amazon.com/dp/0134190440",
			"http://www.informit.com/store/go-programming-language-9780134190440",
			"http://www.barnesandnoble.com/w/1121601944",
			"https://www.gopl.io/ch1.pdf",
			"https://www.gopl.io/ch1.pdf",
			"https://www.gopl.io/ch1.pdf",
			"https://www.gopl.io/ch1.pdf",
			"https://github.com/adonovan/gopl.io/",
			"https://www.gopl.io/reviews.html",
			"https://www.gopl.io/translations.html",
			"https://www.gopl.io/errata.html",
			"http://golang.org/s/oracle-user-manual",
			"http://golang.org/lib/godoc/analysis/help.html",
			"https://github.com/golang/tools/blob/master/refactor/eg/eg.go",
			"https://github.com/golang/tools/blob/master/refactor/rename/rename.go",
			"http://www.amazon.com/dp/0131103628?tracking_id=disfordig-20",
			"http://www.amazon.com/dp/020161586X?tracking_id=disfordig-20",
		}},
		{"https://www.gopl.hogefugapiyo/", true, []string{}},
	}
	for _, test := range tests {
		list, err := Extract(test.url)
		if test.isErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.list, list)
		}
	}
}
