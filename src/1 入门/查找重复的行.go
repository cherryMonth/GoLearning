package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++
	}
}

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() && input.Text() != "end" {
		counts[input.Text()] ++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	f, err := os.Open(os.Args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}

	fileCounts := make(map[string]int)
	countLines(f, fileCounts)
	for line, n := range fileCounts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

/*
单元小结

make关键字可以用来创建动态数组，map是数据结构的类型，在此样例中counts代表一个字典类型，key是字符串类型，value是整形。
input.Scan() 用于从用户输入中获取值
counts[input.Text()]++：数组创建之后，value默认赋值是0。所以当字符串重复时，该值就会加1
map不含某个键时，则自动添加该k-v，v为数据类型的默认值


*/
