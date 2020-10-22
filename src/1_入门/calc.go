package main

import (
	"1_入门/simple"
	"fmt"
	"strconv"
)

func main() {

	v1, err := strconv.Atoi("00123")
	if err != nil {
		fmt.Print("error")
	}

	expr := "sqrt"

	switch expr {
	case "sqrt":
		fmt.Println(simple.Add(5, 4))
	}

	fmt.Println(v1)
}
