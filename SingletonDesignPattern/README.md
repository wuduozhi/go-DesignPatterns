# 单例模式

在 golang 中实现单例模式，转换一下，就是如何保证只创建一个实例的问题。

## 饿汉模式

借用 golang 中的 `init` 函数，在文件加载时，自动执行 init 函数，完成实例创建，避免多线程问题。

## double check

```golang
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
```

## 使用 sync.Once

```golang
var once sync.Once
func GetExampleByOnce()*example{
	once.Do(func() {
		instance = new(example)
	})
	return instance
}
```