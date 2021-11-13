package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	content, err := ioutil.ReadFile("fixture/expected.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := string(content)

	f, _ := os.Open("fixture/go.html")
	defer f.Close()
	doc, _ := html.Parse(f)
	actual := extractStdout(t, forEachNode, doc, start, end)

	assert.Equal(t, expected, actual)
}

func extractStdout(t *testing.T, f func(*html.Node, func(*html.Node), func(*html.Node)), n *html.Node, pre, post func(*html.Node)) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()
	// パイプの作成(r: Reader, w: Writer)
	r, w, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = w
	// テスト対象の関数を実行する
	f(n, pre, post)
	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	w.Close()
	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	// 文字列を取得する
	return strings.TrimRight(buf.String(), "\n")
}
