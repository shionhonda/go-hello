package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	b, _ := ioutil.ReadFile("fixture/sample.xml")
	xmlText := string(b)
	n, err := parse(strings.NewReader(xmlText))
	if err != nil {
		log.Fatal(err)
	}
	visit(n, 0)
}

func visit(n Node, depth int) {
	switch n := n.(type) {
	case CharData:
		fmt.Printf("%*s%s\n", depth*2, "", n)
	case *Element:
		attrs := ""
		for _, a := range n.Attr {
			attrs += fmt.Sprintf(" %s=\"%s\"", a.Name.Local, a.Value)
		}
		if attrs != "" {
			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Type.Local, attrs)
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Type.Local)
		}
		for _, c := range n.Children {
			visit(c, depth+1)
		}
	default:
		log.Fatalf("unknown node type %T", n)
	}
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = e
			} else {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, e)
			}
			stack = append(stack, e) //push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) == 0 {
				panic("xmlselect: no root element")
			}
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, CharData(tok))
		}
	}
	return root, nil
}
