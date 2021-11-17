package counter

import "strings"

type WordCounter int

type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	*w += WordCounter(countWords(string(p)))
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	*l += LineCounter(countLines(string(p)))
	return len(p), nil
}

func countWords(s string) int {
	return len(strings.Fields(s))
}

func countLines(s string) int {
	return strings.Count(s, "\n") + 1
}
