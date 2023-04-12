package main

import "fmt"

type IStack interface {
	empilhar(value int)
	topo(value int)
	desempilhar()
	vazia()
	tamanho()
}

type Stack struct {
	values []int
}

func (s *Stack) empilhar(value int) {
	s.values = append(s.values, value)
}

func (s *Stack) topo() int {
	index := len(s.values) - 1
	return s.values[index]
}

func (s *Stack) desempilhar() int {
	if len(s.values) == 0 {
		return 0
	}
	index := len(s.values) - 1
	valor := (s.values)[index]
	/* A sintaxe s.values[:index] é usada para criar uma nova fatia que contém todos os elementos da fatia original até o índice index (exclusivo),
	ou seja, a nova fatia não inclui o último elemento. */
	s.values = s.values[:index]
	return valor

}

func (s *Stack) vazia() bool {
	if len(s.values) == 0 {
		return true
	}
	return false
}

func (s *Stack) tamanho() int {
	return len(s.values)
}

func main() {
	stack := Stack{}

	stack.empilhar(5)
	stack.empilhar(9)
	stack.empilhar(25)
	fmt.Println(stack.topo())
	fmt.Println(stack.desempilhar())
	fmt.Println(stack.desempilhar())
	fmt.Println(stack.vazia())
	fmt.Println(stack.tamanho())
	fmt.Println(stack.topo())
	fmt.Println(stack.desempilhar())
	fmt.Println(stack.vazia())

}
