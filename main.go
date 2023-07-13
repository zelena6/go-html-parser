package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Tag(f *os.File, s *Stack) {
	closing := false
	char := make([]byte, 1)
	t := make([]byte, 0)

	for {
		_, err := f.Read(char)
		if err != nil {
			if err == io.EOF {
				break

			} else {
				panic(err)
			}
		}

		if string(char) == "/" {
			closing = true
		}

		if string(char) == ">" {
			ast := &AST{
				Tag:    string(t),
				Type:   "closing",
				Childs: make([]*AST, 0),
			}

			if !closing {
				fmt.Println("opening", string(t))
				ast.Type = "opening"
				fmt.Println("opening-parent", string(t), s.Read().Tag)
				s.Read().Childs = append(s.Read().Childs, ast)

				s.Add(ast)
				return
			}

			fmt.Println("closing", string(t))
			s.Pop()
			if s.Read() != nil {
				fmt.Println("closing-parent", string(t), s.Read().Tag)
				s.Read().Childs = append(s.Read().Childs, ast)

			}

			return
		}

		t = append(t, char[0])
	}
}

type AST struct {
	Tag    string `json:"tag"`
	Type   string `json:"type"`
	Childs []*AST `json:"childs"`
}

func main() {
	s := NewStack()
	globalAst := AST{Childs: make([]*AST, 0), Tag: "root", Type: "opening"}
	s.Add(&globalAst)

	f, err := os.Open("index.html")
	if err != nil {
		panic(err)
	}

	char := make([]byte, 1)

	for {
		_, err = f.Read(char)
		if err != nil {
			if err == io.EOF {
				break

			} else {
				panic(err)
			}
		}

		if string(char) == "<" {
			Tag(f, s)
		}
	}

	ReadAST(&globalAst)

	b, _ := json.Marshal(globalAst)
	fmt.Println(string(b))
}

func ReadAST(ast *AST) {
	if len(ast.Childs) == 0 {
		return
	} else {
		for _, a := range ast.Childs {
			fmt.Println(a.Tag, a.Type)
			ReadAST(a)
		}
	}
}
