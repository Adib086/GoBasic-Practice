package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["en"] = "English"
	languages["fr"] = "French"
	languages["es"] = "Spanish"
	languages["de"] = "Germany"
	a := "adib"
	fmt.Println(languages)
	fmt.Println(languages["en"])

	delete(languages, "fr")
	fmt.Println(languages)
	fmt.Println(languages["en"] + " " + a)

	for k, v := range languages {
		fmt.Println("for the key " + k + " for the value " + v)
	}
}
