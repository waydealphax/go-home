package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Faxs = true
	fmt.Println("beans", Faxs)
}
