package main

type Ilist interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}

// Node1P é uma struct que contém o valor do nó e um ponteiro para o próximo nó na lista.
type Node1P struct {
	value int
	next  *Node1P
}

// LinkedList é uma struct que contém um ponteiro para o primeiro nó na lista..
type LinkedList struct {
	head *Node1P
	size int
}

func (linkedlist *LinkedList) Add(value int) {
	newNode := Node1P{value, nil}

	aux := linkedlist.head
	prev := aux

	for aux != nil {
		prev = aux
		aux = aux.next
	}

	if prev == nil {
		linkedlist.head = &newNode
	} else {
		prev.next = &newNode
	}
	linkedlist.size++

}
