package main

type Stack struct {
	List []*AST
}

func NewStack() *Stack {
	return &Stack{
		List: make([]*AST, 0),
	}
}

func (s *Stack) Pop() {
	// fmt.Println("pop", s.List[len(s.List)-1].Tag)
	s.List = s.List[:len(s.List)-1]
}

func (s *Stack) Add(ast *AST) {
	// fmt.Println("add", ast.Tag)
	s.List = append(s.List, ast)
}

func (s *Stack) Read() *AST {
	if len(s.List) == 0 {
		return nil
	}
	return s.List[len(s.List)-1]
}
