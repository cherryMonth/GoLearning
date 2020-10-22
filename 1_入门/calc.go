package main

import (
	"fmt"
	"strconv"
	"入门/1_入门/simple"
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

/*
本节主要是熟悉了Go的引用机制，详情请参考博客 https://www.webook.mobi/display/347
*/
