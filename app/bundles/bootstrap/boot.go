package bootstrap

/**
init 顺序
1、在同一个 package 中，可以多个文件中定义 init 方法

2、在同一个 go 文件中，可以重复定义 init 方法

3、在同一个 package 中，不同文件中的 init 方法的执行按照文件名先后执行各个文件中的 init 方法

4、在同一个文件中的多个 init 方法，按照在代码中编写的顺序依次执行不同的 init 方法

5、对于不同的 package，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 init() 函数

6、如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化，例如：导入顺序 main –> A –> B –> C，则初始化顺序为 C –> B –> A –> main，一次执行对应的 init 方法。

所有 init 函数都在同⼀个 goroutine 内执⾏。
所有 init 函数结束后才会执⾏ main.main 函数。

所以理论上无需关心初始化的依赖顺序 。
但是在某些场景下是会出现可编译依赖问题
例如
boot->b->c
boot->c
boot->d->c

d 调用c 没有问题
b 依赖c 也米有问题
但是 b 预设 c ，d 依赖 b 预设后的 c
此时就出现问题了。

所以不要进行设置操作。而要进行依赖操作

所以 b 依赖 c ，而不是设置 c。


*/

var initList []func()

func AddDInit(init func()) {
	initList = append(initList, init)
}

func StartInit() {
	for _, item := range initList {
		item()
	}
}

func Run() {
	StartInit()
}
