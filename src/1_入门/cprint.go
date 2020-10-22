package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("hello world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}

/*
一点素心

首先这里要用stdio.h和stdlib.h两个头文件，其中puts是stdio.h中的，free是stdlib.h中的
C这个模块只能单独引用，并且要死第一个导入的，不能使用()与其他模块一起引用，否则会报错
导入C模块的语句与 * / 之间不能有换行或者其他字符串
*/
