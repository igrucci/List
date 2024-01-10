package main

import "fmt"

type List[T any] struct {
	tail, head *element[T]
}
type element[T any] struct {
	next  *element[T]
	value T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var all_el []T
	for e := lst.head; e != nil; e = e.next {
		all_el = append(all_el, e.value)
	}
	return all_el
}

func (lst *List[T]) Set(index int, new_v T) {
	cur := lst.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}
	if cur != nil {
		cur.value = new_v
	} else {
		fmt.Println("Index out of range")
	}
}

func main() {
	lst := List[int]{}
	lst.Push(4)
	lst.Push(1)
	lst.Push(6)

	arr := lst.GetAll()
	fmt.Print(arr)

	lst.Set(3, 7)

	arr_after_set := lst.GetAll()
	fmt.Print(arr_after_set)

}
