package main

import (
	"fmt"
)

func main() {
	// const a = 324
	// fmt.Println(a)
	// fmt.Printf("Types: %T\n", a)

	// slice1 := make([]string, 0)
	// slice1 := []string{}
	// fmt.Print(len(slice1), " ", cap(slice1), "\n")
	// slice1 = append(slice1, "hello", "world", "golang", "slices")
	// fmt.Println(slice1)
	// fmt.Print(len(slice1), " ", cap(slice1), "\n")

	// arr := []int{1, 23, 4}
	// // arr = append(arr, 4)
	// fmt.Println(arr)

	// map1 := map[string]int{}
	// for i := 0; i < 5; i++ {
	// 	// map1[string(rune('a'+i))] = i
	// 	// convert i to string and use it as key
	// 	map1[strconv.Itoa(i)] = i
	// }
	// fmt.Println(map1)

	// pointer
	// var a int = 10
	// fmt.Println("value of a is", a)
	// fmt.Println("address of a is", &a)
	// fmt.Printf("type of &a is %T\n", &a)

	// var b *int = &a
	// fmt.Println("value of b is", b)
	// fmt.Println("address of b is", &b)
	// fmt.Printf("type of &b is %T\n", &b)

	// var c **int = &b
	// fmt.Println("value of c is", c)
	// fmt.Println("address of c is", &c)
	// fmt.Printf("type of &c is %T\n", &c)

	// var d ***int = &c
	// fmt.Println("value of d is", d)
	// fmt.Println("address of d is", &d)
	// fmt.Printf("type of &d is %T\n", &d)
	// var e ****int = &d
	// fmt.Println("value of e is", e)
	// fmt.Println("address of e is", &e)
	// fmt.Printf("type of &e is %T\n", &e)

	// a := 10
	// fmt.Println(&a)
	// increment(&a)
	// fmt.Println(a)

	// a := new(int)
	// fmt.Println(a)
	// fmt.Println(*a)

	// var a *int
	// fmt.Println(a)
	// *a = 10
	// fmt.Println(&a)

	dog := Dog{name: "Buddy", age: 3, color: "Brown"}
	cat := Cat{name: "Whiskers", age: 2, color: "Gray"}

	// var walker Walker

	// walker = dog
	// walker.walk()

	// walker = cat
	// walker.walk()

	MakeWalker(dog)
	MakeWalker(cat)

	// GO linked list
	LinkedListTest()
}

type Dog struct {
	name  string
	age   int
	color string
}

type Cat struct {
	name  string
	age   int
	color string
}

func MakeWalker(w Walker) {
	w.walk()
}

type Walker interface {
	walk()
}

func (d Dog) walk() {
	fmt.Println("Dog is walking")
}

func (c Cat) walk() {
	fmt.Println("Cat is walking")
}

func increment(a *int) {
	*a++
	fmt.Println(a)
	fmt.Println(*a)
}
