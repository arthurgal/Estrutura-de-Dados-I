package main

import (
	"errors"
	"fmt"
)

type Ilist interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}

type ArrayList struct {
	values []int
	size   int
}

// Aqui eu comparo se o tamnho que eu quero criar um array é maior que 0
// Se for maior que zero, eu crio o array
func (arraylist *ArrayList) Init(size int) error {
	if size > 0 {
		arraylist.values = make([]int, size)
		arraylist.size = size
		return nil
	} else {
		return errors.New("Você não pode inicializar um arrayList de tamanho 0")
	}
}

// Aqui é a função que dobra o tamanho do array e preenche com os valores que já estavam no array
func (arraylist *ArrayList) Double() {
	doubleValues := make([]int, len(arraylist.values)*2)
	for i := 0; i < len(arraylist.values); i++ {
		doubleValues[i] = arraylist.values[i]
	}
	arraylist.values = doubleValues
}

// Aqui eu adiciono um elemento na ultima posição do array, mas antes comparo se tem espaço no array, senão tiver, chamo a função double()
func (arraylist *ArrayList) Add(value int) {
	if arraylist.size == len(arraylist.values) {
		arraylist.Double()
	}
	arraylist.values[arraylist.size] = value
	arraylist.size++
}

func (arraylist *ArrayList) AddOnIndex(value int, index int) error {
	//O indice onde eu quero adicionar faz parte do meu array?
	if index >= 0 && index <= arraylist.size {

		//preciso aumentar o array?
		if arraylist.size == len(arraylist.values) {
			arraylist.Double()
		}
		//libero o espaço no index desejado
		for i := arraylist.size; i > index; i-- {
			arraylist.values[i] = arraylist.values[i-1]
		}
		//adiciono o valor no index desejado
		arraylist.values[index] = value
		//informo que a quantidade do array aumentou
		arraylist.size++
		return nil
	} else {
		if index < 0 {
			return errors.New("não pode inserir em um index menor que zero")
		} else {
			return errors.New("não pode adicionar um index > arraylist.size")
		}

	}
}

func (arraylist *ArrayList) RemoveOnIndex(index int) error {
	//O indice onde eu quero adicionar faz parte do meu array?
	if index >= 0 && index < arraylist.size {
		//libero o espaço no index desejado
		for i := index; i < arraylist.size-1; i++ {
			arraylist.values[i] = arraylist.values[i+1]
		}
		//informo que a quantidade do array diminuiu
		arraylist.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("não pode remover em um index menor que zero")
		} else {
			return errors.New("não pode remover em um index > arraylist.size")
		}
	}
}

func (arraylist *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index <= arraylist.size {
		return arraylist.values[index], nil
	} else {
		if index < 0 {
			return index, errors.New("Can't get value from arraylist on index < 0")
		} else {
			return index, errors.New("Can't get value from arraylist on index >= arraylist.size")
		}
	}
}

func (arraylist *ArrayList) Set(index int, value int) error {
	if index >= 0 && index <= arraylist.size {
		arraylist.values[index] = value
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't set value on arraylist on index < 0")
		} else {
			return errors.New("Can't set value on arraylist on index >= arraylist.size")
		}
	}
}

func (arraylist *ArrayList) Size() int {
	return arraylist.size
}

func main() {

	fmt.Println("Array List")

	arraylist := ArrayList{}
	tamanho := 10

	arraylist.Init(tamanho)

	for i := 0; i < tamanho; i++ {
		arraylist.values[i] = i
	}

	fmt.Println(arraylist)

	arraylist.AddOnIndex(30, 2)

	fmt.Println(arraylist)

	arraylist.AddOnIndex(30, 5)

	fmt.Println(arraylist)

	arraylist.RemoveOnIndex(2)

	fmt.Println(arraylist)

	fmt.Println("#######################")

	fmt.Println("Linked List")
}
