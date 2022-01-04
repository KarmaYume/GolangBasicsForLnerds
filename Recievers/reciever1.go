//basically associating a function with a struct using a variable to the pointer of that struct

package main

import "fmt"

type something struct {
	MyfavAnime string
}

func (s *something) printAnimeName() string {
	return s.MyfavAnime
}

func main() {

	x1 := something{
		MyfavAnime: "Gintama",
	}

	x2 := something{
		MyfavAnime: "Kaguya Sama Love Is War",
	}

	fmt.Println("My first favorite anime is : ", x1.printAnimeName())
	fmt.Println("My second favorite anime is : ", x2.printAnimeName())

}
