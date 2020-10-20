package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name           string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("i am fly")
}

func main() {
	sparrow := &Bird{"sss", 2}
	s := reflect.ValueOf(sparrow).Elem()
	typeOf := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d %s %s = %v\n", i, typeOf.Field(i).Name, f.Type(), f.Interface())
	}
}

/*
本周小结
本节使用了go的反射技术读取了定义实例的属性和默认值，go与C++类似，可以使用&取地址
*/
