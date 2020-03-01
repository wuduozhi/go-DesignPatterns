package SingletonDesignPattern

import (
	"sync"
)

type example struct {
	name string
}

var instance *example

/***
饿汉模式
使用 init ,文件加载时自动执行 init 函数，创建实例
func init(){
	instance = new(example)
}

func GetExample()*example{
	return instance
}
***/

var mux sync.Mutex
func GetExampleByDoubleCheck()*example{
	//double check
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = &example{}
		}
	}
	return instance
}

/*
Mutex由操作系统实现，而atomic包中的原子操作则由底层硬件直接提供支持。
在 CPU 实现的指令集里，有一些指令被封装进了atomic包，这些指令在执行的过程中是不允许中断（interrupt）的，
因此原子操作可以在lock-free的情况下保证并发安全，并且它的性能也能做到随 CPU 个数的增多而线性扩展。
 */

var once sync.Once
func GetExampleByOnce()*example{
	once.Do(func() {
		instance = new(example)
	})
	return instance
}