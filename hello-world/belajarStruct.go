package main

import "fmt"

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name:", s1.name)
	fmt.Println("grade:", s1.grade)
}

type student struct {
	name  string
	grade int
}
