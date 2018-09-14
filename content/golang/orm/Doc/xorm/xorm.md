# ![](./favicon.png)[XORM使用整理](http://www.xorm.io/)
## 简介
xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。xorm的目标并不是让你完全不去学习SQL，我们认为SQL并不会为ORM所替代，但是ORM将可以解决绝大部分的简单SQL需求。xorm支持两种风格的混用。

特性

* 支持Struct和数据库表之间的灵活映射，并支持自动同步表结构
* 事务支持
* 支持原始SQL语句和ORM操作的混合执行
* 使用连写来简化调用
* 支持使用Id, In, Where, Limit, Join, Having, Table, Sql, Cols等函数和结构体等方式作为条件
* 支持级联加载Struct
* 支持LRU缓存(支持memory, memcache, leveldb, redis缓存Store) 和 Redis缓存
* 支持反转，即根据数据库自动生成xorm的结构体
* 支持事件
* 支持created, updated, deleted和version记录版本（即乐观锁）

## 初始化

	engine, err := xorm.NewEngine("mysql", "root:1234@tcp(127.0.0.1:3306)/operation_i_dzz?charset=utf8")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

* 每个数据库都需要单独创建一个数据库引擎对象，所以需要在代码中全局保存这个数据库操作对象
* xorm每次操作是基于会话操作的，所以虽然它是针对每个数据库都是单独创建的数据库引擎对象，但仍然是协程安全的。

### 查

	item := make([]*Tt, 0)

	if errMsg := engine.Where("Id=?", 1).Limit(1, 0).Find(&item); errMsg != nil {
		fmt.Printf(errMsg.Error())
		return
	}

	fmt.Printf("name:%v", item[0].Name)

* 使用Get, Find, Count, Rows, Iterate作为最终结果的获取

	
### 增

	if count, errMsg := engine.Insert(&Tt{
		Id:   3,
		Name: "h3",
	}); errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	} else {
		fmt.Println("数据insert成功，影响记录数:", count)
	}

* 如果传入的是一个[]*Tt，则会进行批量插入
* 批量插入会自动生成Insert into table values (),(),()的语句，因此各个数据库对SQL语句有长度限制，因此这样的语句有一个最大的记录数，根据经验测算在150条左右。大于150条后，生成的sql语句将太长可能导致执行失败。因此在插入大量数据时，目前需要自行分割成每150条插入一次。
* 这里虽然支持同时插入，但这些插入并没有事务关系。因此有可能在中间插入出错后，后面的插入将不会继续。此时前面的插入已经成功，如果需要回滚，请开启事物。

### 改

	item := &Tt{
		Id:   3,
		Name: "h5",
	}

	if count, errMsg := engine.
		Where("Id=?", item.Id).And("Name<>?", item.Name).
		AllCols().Update(item); errMsg != nil { //// 如果不使用AllCols，则Update只会提交非0值
		fmt.Println(errMsg.Error())
		return
	} else {
		fmt.Println("数据update成功，影响记录数:", count)
	}

* 使用Update进行更新，但更新条件需要使用相关查询条件组装函数进行指定更新条件
* 默认情况下，Update只会更新非0值，也就是说，数值0、字符串""，bool值false都不会默认提交到数据库，除非使用AllCols()指定

### 删
	if count, errMsg := engine.Where("Id=?", 4).Delete(&Tt{}); errMsg != nil {
		fmt.Printf(errMsg.Error())
	} else {
		fmt.Println("数据update成功，影响记录数:", count)
	}

* 使用delete进行数据删除
* Delete中的结构体中，如果存在非0值，则会使非0值作为删除条件
* 如果需要指定额外的条件，需要使用相关查询条件语句指定

### 执行原始sql语句

	data := make([]*Tt, 0)

	if errMsg := engine.Sql("select * from Tt").Find(&data); errMsg != nil {
		fmt.Printf(errMsg.Error())
		return
	}

	for _, item := range data {
		fmt.Printf("\r\nId:%v	name:%v", item.Id, item.Name)
	}

* 使用Sql、Exec进行原始sql语句的执行

### 事务操作
	session := engine.NewSession()
	defer session.Close()

	isOk := false
	defer func() {
		if isOk {
			session.Commit()
		} else {
			session.Close()
		}
	}()

	if _, errMsg := session.Insert(&Tt{
		Id:   4,
		Name: "h4",
	}); errMsg != nil {
		fmt.Println("插入异常：", errMsg.Error())
		isOk = false
	}

	if _, errMsg := session.Where("Id=?", 3).AllCols().Update(&Tt{
		Id:   3,
		Name: "hh3",
	}); errMsg != nil {
		fmt.Println("更新异常：", errMsg.Error())
		isOk = false
	}

* 事务操作需要显式地单独开启一个会话

## 总结

**优点**

1. 分离出来的子项目：[core](https://github.com/go-xorm/core)能够很方便地对数据库表结构进行解析和数据库类型到go类型的转换
2. 整体来说使用也确实很简单
3. 支持对数据库表结构和数据的导入导出操作
4. 支持读取数据表的结构信息，并提供有类型映射
5. 有对应的命令行代码生成工具，提供数据表到代码的映射，使用tpl作为代码生成模版

**缺点**

1. 源码基本没啥注释，很不方便后期维护
2. 数据删除接口Delete使用很不方便，必须要传一个结构体，而删除条件另外指定更实在一点
3. 无基于协程级别的事务操作
4. 需要为每个数据库连接实例单独定义变量，比较麻烦（也是可以接受的）
5. 更新时，如果不指定AllCols()，则只会更新非0值的字段

**整体来说，没有beego好用；但用来写数据库相关工具一定是上上之选，因为有较完整的数据库结构的读写支持**