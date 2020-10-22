package main

import "fmt"

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, values := range values {
		sum += values
	}
	resultChan <- sum
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 1; i++ {
		resultChan := make(chan int, 2)

		go sum(values[:len(values)/2], resultChan)
		go sum(values[len(values)/2:], resultChan)

		sum1, sum2 := <-resultChan, <-resultChan
		fmt.Println("Result ", sum1, sum2, sum1+sum2)
	}

}

/*
本节小结
Go语言使用了channel（管道）来实现CSP模型。
*/
