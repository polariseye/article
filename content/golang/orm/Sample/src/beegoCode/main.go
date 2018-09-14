package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func initDb() {
	// 注册驱动，默认已经注册常见的mysql、sqlit和Postgres，所以一般可以不用注册
	if errMsg := orm.RegisterDriver("mysql", orm.DRMySQL); errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 注册数据库连接,必须要注册一个名为default的连接，当作默认使用
	if errMsg := orm.RegisterDataBase("default", "mysql",
		"root:1234@tcp(127.0.0.1:3306)/operation_i_dzz?charset=utf8&parseTime=true"); errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 注册model类型，需要注意的是，结构名必须和表名基本一致（可以比表名少前缀或后缀）
	orm.RegisterModel(new(Tt))
	// orm.RegisterModelWithPrefix("p_", new(Tt))
	//orm.RegisterModelWithSuffix("")
}

type Tt struct {
	Id   int
	Name string
}

func main() {
	initDb()
	deleteTst()
}

// 基于原始sql查询
func rawQueryTst() {
	// 创建Orm对象
	o := orm.NewOrm()
	// 执行sql
	result := o.Raw("select Id,Name from tt limit 10") // 执行sql语句

	var ( // 定义用于接收返回值的slice
		idList   []int
		nameList []string
	)

	// 获取结果
	count, errMsg := result.QueryRows(&idList, &nameList)
	if errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 打印结果
	fmt.Println("查询到记录数：", count)
	for index, _ := range idList {
		fmt.Printf("\r\nPartnerId:%v	UserId:%v	LoginDate:%v", idList[index], nameList[index], 1)
	}
}

// 使用model形式查询
func modelQueryTst1() {
	o := orm.NewOrm()
	o.Using("default")

	var data []Tt
	count, errMsg := o.QueryTable("tt").All(&data)
	if errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	fmt.Println("查询到记录数：", count)
	for _, item := range data {
		fmt.Printf("\r\nPartnerId:%v	UserId:%v	", item.Id, item.Name)
	}
}

// 使用model形式查询
func modelQueryTst2() {
	// 新建一个操作对象
	o := orm.NewOrm()
	// 指定使用的数据库连接，默认为default
	o.Using("default")

	var data []Tt //// 定义接收数据的slice

	query := o.QueryTable("tt")        //// 使用QueryTable指定要查询的表
	query = query.Filter("id__gte", 2) //// 添加筛选条件
	count, errMsg := query.All(&data)  //// All表示接收所有数据
	if errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 打印结果
	fmt.Println("查询到记录数：", count)
	for _, item := range data {
		fmt.Printf("\r\nId:%v	Name:%v	", item.Id, item.Name)
	}
}

// 使用model形式查询
func modelQueryTst3() {
	// 新建一个操作对象
	o := orm.NewOrm()
	// 指定使用的数据库连接，默认为default
	o.Using("default")

	data := Tt{}                  //// 定义接收数据的slice
	data.Id = 1                   //// 指定查询条件
	errMsg := o.Read(&data, "Id") //// 指定使用哪一列进行查询
	if errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 打印结果
	fmt.Printf("查询到的值：name:%v", data.Name)
}

// 数据插入
func insertTst() {
	// 新建一个操作对象
	o := orm.NewOrm()
	// 指定使用的数据库连接，默认为default
	o.Using("default")

	data := Tt{
		Id:   3,
		Name: "h3",
	}

	count, errMsg := o.Insert(&data)
	if errMsg != nil {
		fmt.Printf("执行insert异常，错误信息:%v", errMsg.Error())
		return
	}

	fmt.Printf("影响的记录行数：%v", count)
}

// 更新数据
func updateTst() {
	// 新建一个操作对象
	o := orm.NewOrm()
	// 指定使用的数据库连接，默认为default
	o.Using("default")

	data := Tt{
		Id:   1,
		Name: "hh1",
	}

	count, errMsg := o.Update(&data, "Name")
	if errMsg != nil {
		fmt.Printf("执行update异常，错误信息:%v", errMsg.Error())
		return
	}

	fmt.Printf("影响的记录行数：%v", count)
}

// 删除数据
func deleteTst() {
	// 新建一个操作对象
	o := orm.NewOrm()
	// 指定使用的数据库连接，默认为default
	o.Using("default")

	data := Tt{
		Id:   3,
		Name: "hh1",
	}

	count, errMsg := o.Delete(&data, "Id")
	if errMsg != nil {
		fmt.Printf("执行Delete异常，错误信息:%v", errMsg.Error())
		return
	}

	fmt.Printf("影响的记录行数：%v", count)
}

// 事务测试
func TransactionTst() {
	// 新建一个操作对象
	o := orm.NewOrm()
	// 指定使用的数据库连接，默认为default
	o.Using("default")

	// 开启事务
	if errMsg := o.Begin(); errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 事务结果操作
	isOk := false
	defer func() {
		if isOk {
			o.Commit() // 提交事务
		} else {
			o.Rollback() // 回滚事务
		}
	}()

	// 插入记录
	if _, errMsg := o.Insert(&Tt{
		Id:   3,
		Name: "h3",
	}); errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 更新数据
	if _, errMsg := o.Update(&Tt{
		Id:   1,
		Name: "h1",
	}); errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	}

	// 设置处理成功
	isOk = true
}
