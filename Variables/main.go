package main

import "fmt"

const pass = "123456"

func main() {
	var userName string = "Adib"
	fmt.Println(userName)
	fmt.Printf("the type of var is: %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of var is: %T ", isLoggedIn)

	var smallValue uint8 = 244
	fmt.Println(smallValue)
	var intValue int = 244
	fmt.Println(intValue)

	var floatValue float32 = 244.25632154
	fmt.Println(floatValue)

	var floatValue2 float64 = 244.25632154
	fmt.Println(floatValue2)

	var defaultint int
	fmt.Println(defaultint)

	var website = "Google.com"
	fmt.Println(website)

	name := "adib"
	fmt.Println(name)

	fmt.Println(pass)

}
