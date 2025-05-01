package main

import (
	"fmt"
	"sort"
)

func main() {
	var list = []string{"A", "B", "C"}

	fmt.Println(list)

	list = append(list, "D", "E", "F")

	fmt.Println(list)

	list = append(list[1:3])
	fmt.Println(list)

	highScore := make([]int, 4)

	highScore[0] = 1
	highScore[1] = 2
	highScore[2] = 3
	highScore[3] = 4

	fmt.Println(highScore)

	highScore = append(highScore, 5)
	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"Math", "English", "Chemistry", "Bangla"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
