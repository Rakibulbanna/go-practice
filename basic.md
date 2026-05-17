## Go Basic

```go
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ==========================================
// 1. Variables and Basic Types
// ==========================================
func variablesDemo() {
	fmt.Println("--- 1. Variables ---")

	// 'var' keyword for explicit declaration
	var i int = 10
	// int32, int64, float32, float64, bool, string
	var i32 int32 = 10
	var i64 int64 = 10
	var f32 float32 = 3.14
	var f64 float64 = 3.14
	var b bool = true
	var s string = "Hello Go"

	// Type inferred declaration
	var inferred = 42

	// Short variable declaration (most common inside functions)
	// := creates and assigns the variable
	name := "Alice"
	age := 30

	// Constants
	const pi float32 = 3.14159

	fmt.Printf("Types: %T, %T, %T, %T, %T, %T, %T, %T\n", i, i32, i64, f32, f64, b, s, inferred)
	fmt.Println("Short declaration:", name, age, inferred)
	fmt.Printf("%T %T %T\n", age, inferred, pi)
}

// ==========================================
// 2. Control Structures (If, For, Switch)
// ==========================================
func controlStructuresDemo() {
	fmt.Println("\n--- 2. Control Structures ---")

	// If statement (can have a short assignment before the condition)
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	}

	// Go only has one looping construct: the 'for' loop
	// Basic for loop
	fmt.Print("For loop: ")
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// "While" loop equivalent (just a for with condition)
	count := 0
	for count < 2 {
		count++
	}

	// Switch statement (no 'break' needed in Go, it breaks automatically)
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	default:
		fmt.Println("Other day")
	}
}

// ==========================================
// 3. Arrays, Slices, and Maps
// ==========================================
func collectionsDemo() {
	fmt.Println("\n--- 3. Arrays, Slices, Maps ---")

	// Array: fixed size
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice: dynamic size (built on top of arrays) - much more common
	slice := []int{10, 20, 30}
	slice = append(slice, 40) // Adding to a slice
	fmt.Println("Slice:", slice)

	// Map: key-value pairs (like dictionaries in Python or Objects in JS)
	userAges := make(map[string]int)
	userAges["Alice"] = 28
	userAges["Bob"] = 35
	fmt.Println("Map:", userAges)

	// make dynamic type as like as js oject user name, age, email
	user := make(map[string]interface {
	})
	user["name"] = "Alice"
	user["age"] = 28
	user["email"] = "[EMAIL_ADDRESS]"
	fmt.Println("Dynamic Map:", user, "%T", user["name"])

	// Iterating over a map or slice using 'range'
	for key, value := range userAges {
		fmt.Printf("%s is %d years old\n", key, value)
	}
}

// ==========================================
// 4. Functions & Pointers
// ==========================================

// Function returning multiple values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func pointersDemo() {
	fmt.Println("\n--- 4. Pointers ---")
	// Go has pointers but NO pointer arithmetic
	val := 10
	ptr := &val // & gets the memory address of val

	fmt.Println("Value:", val)
	fmt.Println("Pointer address:", ptr)
	fmt.Println("Dereferenced Pointer Value:", *ptr) // * accesses the value at the address

	// Modifying value through pointer
	*ptr = 20
	fmt.Println("Modified Value:", val)
}

// ==========================================
// 5. Structs, Methods, and Interfaces
// ==========================================

// Struct: Custom data type grouping related variables
type Person struct {
	Name string
	Age  int
}

// Method: A function attached to a specific type (Person in this case)
func (p Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d\n", p.Name, p.Age)
}

// Interface: Defines a behavior (a set of method signatures)
type Speaker interface {
	Speak() string
}

// Dog implements the Speaker interface because it has a Speak() method
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func OOPDemo() {
	fmt.Println("\n--- 5. Structs & Interfaces ---")
	p := Person{Name: "Charlie", Age: 40}
	p.Greet()

	var s Speaker = Dog{}
	fmt.Println("Dog says:", s.Speak())
}

// ==========================================
// 6. Goroutines (Concurrency)
// ==========================================
func printNumbers(prefix string, wg *sync.WaitGroup) {
	// Defer executes when the function finishes
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func goroutineDemo() {
	fmt.Println("\n--- 6. Goroutines ---")
	// WaitGroup is used to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Starting two concurrent goroutines
	wg.Add(2)
	// The 'go' keyword starts a new concurrent thread (goroutine)
	go printNumbers("Routine A", &wg)
	go printNumbers("Routine B", &wg)

	wg.Wait() // Wait for both to finish
	fmt.Println("Goroutines finished!")
}

// ==========================================
// 7. Channels
// ==========================================
func channelDemo() {
	fmt.Println("\n--- 7. Channels ---")

	// Channels are pipes that connect concurrent goroutines
	// make(chan Type)
	ch := make(chan string)

	// Goroutine sending data to the channel
	go func() {
		ch <- "Hello from channel!" // Send value to channel
	}()

	// Main goroutine receiving data from the channel
	msg := <-ch // Receive value from channel (blocks until data is ready)
	fmt.Println("Received:", msg)

	// Buffered Channels: can hold specific number of values without blocking
	bufCh := make(chan int, 2)
	bufCh <- 100
	bufCh <- 200
	// We can pull these out without needing a separate goroutine
	fmt.Println("Buffered 1:", <-bufCh)
	fmt.Println("Buffered 2:", <-bufCh)
}

func main() {
	fmt.Println("=== Welcome to Go Basics ===")

	// variablesDemo()
	// controlStructuresDemo()
	collectionsDemo()
	// pointersDemo()
	// OOPDemo()

	// // Error handling demo
	// fmt.Println("\n--- Error Handling ---")
	// result, err := divide(10, 2)
	// if err != nil { // Standard pattern: check if error is not nil
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("10 / 2 =", result)
	// }

	// goroutineDemo()
	// channelDemo()

	fmt.Println("\n=== End of Go Basics ===")
}
```
