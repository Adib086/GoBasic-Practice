package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var age int
	fmt.Scan(&age)
	fmt.Println(age)

	mark, _ := reader.ReadString('\n')

	sum, err := strconv.ParseFloat(strings.TrimSpace(mark), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum + 1)
	}

}
