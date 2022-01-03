package main

var waifu = "kaguya shinomiya"

func main() {
	waifu := "Komi Shouko"                                   //variable shadowing
	println("My favorite anime waifu is : ", myWaifu(waifu)) //here prefernce is given to local variable over global coz of variable shadowing

}

func myWaifu(waifu string) string { //variable shadowing
	return waifu
}
