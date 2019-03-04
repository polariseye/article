Golang语法原理--defer
------------------------------------
#defer 原理

defer用于注册延时调用方法，在return的前一刻以先进后出的方式调用所有defer函数。
简单如下:

    var lockObj = sync.Mutex{}
    
    func add() {
    	lockObj.Lock()
    	defer lockObj.Unlock()
    
    	i = i + 1
    }

官方原句如下：
> Each time a “defer” statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. If a deferred function value evaluates to nil, execution panics when the function is invoked, not when the “defer” statement is executed.

中文翻译
> 每当defer执行时，函数和函数参数值都将被保存到一个堆栈中，但函数实际上并没有被调用。而是在对应函数返回那一刻以先进后出的方式被调用。如果函数是一个nil指针，则在被调用时panic，而不是在defer定义的地方执行。

具体如下:
1. defer只在函数调用结束之际执行
2. defer仍然可以更改函数调用的返回值
3. defer语句会把函数与相关参数值入栈，所以会有一定的性能损耗，相当于值传递
4. 如果是闭包调用，则会引用外部变量的最终值，相当于引用传递

测试代码1

	package main
	
	import "fmt"
	
	func main() {
		for i := 0; i < 4; i++ {
			defer func() {
				fmt.Println(i)
			}()
		}
	}

**由于是闭包调用defer,所以会直接使用最终的i值**,输出结果为:

	4
	4
	4
	4

测试代码2

	package main
	
	func main() {
		var i = 1
		defer func(val int) {
			println("defer:", val)
		}(i)
	
		i = 2
		println("直接调用:", i)
	}

**defer的调用参数和普通调用一样，参数值在调用时就确认了**。输出结果如下:

	直接调用: 2
	defer: 1

defer可以理解为执行在了return和函数调用结束之间。比如:``return 1``实际被拆分成了

	1. 返回值 = 1
	2. 调用defer函数
	3. 空的return


defer其他说明:
1. 因为defer是在函数结束时执行。所以语句``defer lockObj.Unlock()`` 实际锁住整个了函数。如果使用不当，会容易导致defer执行效率极低

#recover与defer
recover用于捕获panic，以便在出现不可逆错误时，能够做一些善后工作。但使用中有一些具体要求：

1. recover只能用在defer语句中
2. 不能直接``defer recover()``，如果真这么干了，则这句话就相当于没有执行样，不会起任何作用。
3. recover只能出现在defer的函数调用中，不能在嵌套中调用

示例1：
	package main
	
	func main() {
		Hello()
	
		println("调用成功")
	}
	
	func Hello() {
		defer println("退出了")
		defer recover()
	
		panic("add")
	}

这个不会输出``调用成功``，而是直接panic了。如果Hello函数内没有panic,则不会出任何问题,具体输出为：

	退出了
	panic: add'
	
	goroutine 1 [running]:
	main.Hello()
	        E:/GoTestCode/src/testCode/deferTst/main.go:13 +0x9b
	main.main()
	        E:/GoTestCode/src/testCode/deferTst/main.go:4 +0x29
	exit status 2

示例2:

	package main
	
	func main() {
		Hello()
	
		println("调用成功")
	}
	
	func Hello() {
		defer println("退出了")
		defer func() {
			if err := recover(); err != nil {
				println("捕获到panic:", err.(string))
			}
		}()
	
		panic("add")
	}

这个是完全正常的使用，所以最后会输出``调用成功``,具体输出内容为:

	捕获到panic: add
	退出了
	调用成功

示例3：

	package main
	
	func main() {
		Hello()
	
		println("调用成功")
	}
	
	func Hello() {
		defer println("退出了")
		defer func() {
			func() {
				if err := recover(); err != nil {
					println("捕获到panic:", err.(string))
				}
			}()
		}()
	
		panic("add")
	}

由于recover只在直接调用的层次才会起作用，所以这个例子仍然会panic，具体输出为:

	退出了
	panic: add
	
	goroutine 1 [running]:
	main.Hello()
	        E:/GoTestCode/src/testCode/deferTst/main.go:19 +0x91
	main.main()
	        E:/GoTestCode/src/testCode/deferTst/main.go:4 +0x29
	exit status 2