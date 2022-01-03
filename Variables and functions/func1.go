package main

import "fmt"

func main() {

	//here we assign the variable x and y with return values of nice() along with its return types
	x, y := nice()
	fmt.Println(x, y)
}

//here we created a function "nice()" with return types string and int
func nice() (string, int) {
	return "nice", 60 + 9
}
