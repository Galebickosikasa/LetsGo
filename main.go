package main

import (
	"fmt"
)

type Student struct {
	Name     string
	LastName string
	id       int
}

func NewStudent(name string, lastName string, id int) *Student {
	return &Student{name, lastName, id}
}

func test() {
	fmt.Println("Lets Go")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	var a []int
	fmt.Printf("size = %d\n", len(a))
	a = append(a, 228, 1337)
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}
	fmt.Printf("\nsize = %d\n", len(a))
	fmt.Println(a)

	var student = NewStudent("John", "Smith", 23)
	fmt.Println(*student)
}

func main() {
	test()
}
