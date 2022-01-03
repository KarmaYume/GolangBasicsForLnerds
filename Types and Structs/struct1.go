package main

import "fmt"

//struct is basically a user defined varibale
//mainly used to maintain data which can be a mix of diffrent bits of info with different data types

type User struct {
	FirstName         string
	LastName          string
	FavoriteAnime     string
	FavoriteAnimeChar string
}

func main() {
	user1 := User{
		FirstName:         "Karma",
		LastName:          "Yume",
		FavoriteAnime:     "Gintama",
		FavoriteAnimeChar: "Kaguya Shinomia",
	}

	fmt.Println(user1.FavoriteAnime, user1.FavoriteAnimeChar)
}
