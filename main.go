package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else {
		return "tidak lulus"
	}
}

// debug
func main() {
	fmt.Println(GraduateStudent(100, 4)) // Output: "lulus"
	fmt.Println(GraduateStudent(80, 5))  // Output: "tidak lulus"
}
