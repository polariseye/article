想要爱你不容易--golang语法原理之接口
-----------------------------
interface是golang开发常用的数据类型之一，但坑却是不少。本文将详细说明接口原理和接口的常见坑。本文假设读者已了解golang基本语法。

# 接口原理
接口数据类型在底层实现实际占用了两个指针大小的内存（不包含具体元数据信息和具体数据），并以实现了两种方式的接口 不同各类的接口其指针含义也不一样：
1. 空类型接口: 没有函数声明的接口或者普通定义的``a:=interface{}``是空接口类型，它的内存结构是下面的``eface``
2. 非空类型接口: 我们通过关键字``interface``定义的包含有方法的接口是非空类型接口,它的内存结构对应的是下面的``iface``

```
	// 非空接口元数据的内存定义
	type iface struct {
		tab  *itab
		data unsafe.Pointer
	}
	
	// 空接口的内存定义
	type eface struct {
		_type *_type
		data  unsafe.Pointer
	}

	// 接口具体类型的元数据
	type itab struct {
		inter *interfacetype
		_type *_type
		hash  uint32 // copy of _type.hash. Used for type switches.
		_     [4]byte
		fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
	}

	// interface数据类型对应的type
	type interfacetype struct {
		typ     _type
		pkgpath name
		mhdr    []imethod
	}

	// 具体数据类型的元数据
	type _type struct {
	    size       uintptr // type size
	    ptrdata    uintptr // size of memory prefix holding all pointers
	    hash       uint32  // hash of type; avoids computation in hash tables
	    tflag      tflag   // extra type information flags
	    align      uint8   // alignment of variable with this type
	    fieldalign uint8   // alignment of struct field with this type
	    kind       uint8   // enumeration for C
	    alg        *typeAlg  // algorithm table
		// gcdata stores the GC type data for the garbage collector.
		// If the KindGCProg bit is set in kind, gcdata is a GC program.
		// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
		gcdata    *byte // garbage collection data
		str       nameOff  // string form
		ptrToThis typeOff  // type for pointer to this type, may be zero
	}
```

空类型的接口存储了一个指向具体类型的``_type``指针和指向具体数据的``unsafe.Pointer``指针。golang中大部分数据类型都可以抽象出``_type``，它也包含了一个类型的具体元数据信息

非空类型的接口包含一个``itab``指针和一个指向具体数据的``unsafe.Pointer``指针。``itab``包含了这个接口的方法元数据信息，主要信息在``typ``字段中。它包含了包路径，方法列表和其他``_type``相关的元数据信息

说明：
 * 以上代码搜集自标准库的runtime包，golang版本：1.11.5

# 接口测试

	type empty interface {
	}
	
	type Person struct {
		name string
	}
	
	func (this *Person) Name() string {
		return this.name
	}

	func main() {	
		var iVal empty = Person{"asd"}
		var val = (*eface)(unsafe.Pointer(&iVal))
		println("第一次:", val._type.kind == uint8(reflect.Struct))
	
		iVal = "abcd123123123"
		val = (*eface)(unsafe.Pointer(&iVal))
		println("第二次:", val._type.kind == uint8(reflect.String))
	}

这个目的是查找出具体类型信息，这个可比反射快多了。另外，如果使用``iface``来类型转换将会panic。具体输出结果为:

	第一次: true
	第二次: true

一个有趣的测试，main函数替换为:

	func main() {	
			var iVal empty = &Person{"asd"} ////---------这行多了个取地址
			var val = (*eface)(unsafe.Pointer(&iVal))
			println("第一次:", val._type.kind == uint8(reflect.Ptr))
		}

此时不会输出``true``。具体``val._type.kind``的值为56。原因暂时不知。知道的朋友可以留言告知。

# 接口的注意事项
1. 接口是由类型和数据组成，所以``a:= interface { Hello(); }(nil);``，中a不为nil。表达式``a==nil``的值为``false``
2. 一个值类型也可以赋值给接口，接口将保存对应值类型的副本

## 接口中nil问题的处理
在知道interface具体类型的情况下，可以通过 ``var a inteface=typea(b);a.(typea)==nil;``这种方式判断。如果不知道具体类型，又想知道是否为nil。则需要转换到对应的内存类型来判断。具体如下:

	// 判断一个数据是否为nil。包含把一个具体类型赋值给interface的情况
	// 返回值:
	// bool: true：为nil false:不为nil
	func IsNil(val interface{}) bool {
		if val == nil {
			return true
		}
	
		type InterfaceStructure struct {
			pt uintptr // 指向interface方法表的指针
			pv uintptr // 指向对应值的指针
		}
		is := *(*InterfaceStructure)(unsafe.Pointer(&val))
		if is.pv == 0 {
			return true
		}
	
		return false
	}


# 参考文档
[golang interface深度解析](https://blog.csdn.net/D_Guco/article/details/78507999)
[Go语言interface实现原理详解](https://www.jianshu.com/p/70003e0f49d1)
[Go Interface 源码剖析](http://legendtkl.com/2017/07/01/golang-interface-implement/)