package SingletonDesignPattern

import (
	"fmt"
	"testing"
)

func TestGetExampleByDoubleCheck(t *testing.T){
	s := GetExampleByDoubleCheck()
	s.name = "第一次赋值单例模式"
	fmt.Println(s.name)

	s2 := GetExampleByDoubleCheck()
	fmt.Println(s2.name)

	s3 := GetExampleByOnce()
	s3.name = "第二次赋值单例模式"
	fmt.Println(s3.name)

	s4 := GetExampleByOnce()
	fmt.Println(s4.name)
}



