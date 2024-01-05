package main
import "fmt"

type Person struct {
	name string
	age int
}

func sum(nums ...int) (int, bool) {
  total := 0

	for _,num := range nums {

		total += num
	}

	return total, true
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for key := range m {
		r = append(r, key)
	}
	return r
}

type List[T any] struct {
  head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e!= nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func main() {
	sliceA := []int{1, 2, 3, 4}
	total, _ := sum(sliceA...)
	fmt.Println(total)

	var m = map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println("keys", MapKeys(m))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
  fmt.Println("list:", lst.GetAll())
}
