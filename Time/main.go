package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createDate := time.Date(2020, time.October, 10, 3, 5, 0, 0, time.UTC)

	fmt.Println(createDate.Format("2006-01-02 15:04:05 Monday"))
}
