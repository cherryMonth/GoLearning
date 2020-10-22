package simple

import (
	"testing"
)

func TestAdd(t *testing.T) {
	v := Add(5, 4)
	if v != 5+4 {
		t.Error("error ")
	}
}

/*
Go的测试代码
*/
