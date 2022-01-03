//basically creating a reference to a variable
//demonstrating how to update value of variable using a pointer

package main

import (
	"fmt"
)

func main() {
	var favAnime string
	favAnime = "Gintama"
	fmt.Println("Current Favorite Anime : ", favAnime)
	//passing a reference to the variable
	newFavAnime(&favAnime)
	fmt.Println("Updated Favorite Anime : ", favAnime)
}

//creating a new funtion with a parameter which is pointer to a string
func newFavAnime(a *string) {
	fmt.Println("memory address of the pointer", a)
	newAnime := "Kaguya Sama Love is war"
	*a = newAnime
}
