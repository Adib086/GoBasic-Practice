package main

import "fmt"

func main() {
	myCar := Car{
		Name:         "Toyota",
		Engine:       "V6",
		Color:        "Red",
		Sports:       false,
		Registration: 15674,
	}
	fmt.Println(myCar)
}

type Car struct {
	Name         string
	Engine       string
	Color        string
	Sports       bool
	Registration int
}
