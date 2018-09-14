package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type Tt struct {
	Id   int
	Name string
}

func main() {
	// xorm是
	engine, err := xorm.NewEngine("mysql", "root:1234@tcp(127.0.0.1:3306)/operation_i_dzz?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	DeleteTst2(engine)
}

// 打印数据表的基本信息
func OutDbInfo(engine *xorm.Engine) {
	var (
		tb  []*core.Table
		err error
	)
	tb, err = engine.DBMetas()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, tbItem := range tb {
		fmt.Println("表名:", tbItem.Name)
		for _, col := range tbItem.Columns() {
			nowTp := core.SQLType2Type(col.SQLType)
			fmt.Printf("Name:%v DbType:%v goType:%v \r\n ", col.Name, col.SQLType, nowTp.String())
		}
	}
}

// 插入测试
func InsertTst(engine *xorm.Engine) {
	if count, errMsg := engine.Insert(&Tt{
		Id:   3,
		Name: "h3",
	}); errMsg != nil {
		fmt.Println(errMsg.Error())
		return
	} else {
		fmt.Println("数据insert成功，影响记录数:", count)
	}
}

// 更新测试
func Update(engine *xorm.Engine) {
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
}

// 删除测试
func DeleteTst(engine *xorm.Engine) {
	if count, errMsg := engine.Where("Id=?", 4).Delete(&Tt{}); errMsg != nil {
		fmt.Printf(errMsg.Error())
	} else {
		fmt.Println("数据update成功，影响记录数:", count)
	}
}

// 删除测试
func DeleteTst2(engine *xorm.Engine) {
	data := &Tt{
		Id:   3,
		Name: "h4",
	}
	if count, errMsg := engine.Id().Delete(data); errMsg != nil {
		fmt.Printf(errMsg.Error())
	} else {
		fmt.Println("数据Delete成功，影响记录数:", count)
	}
}

// 查询测试
func QueryTst1(engine *xorm.Engine) {
	item := make([]*Tt, 0)

	if errMsg := engine.Where("Id=?", 1).Limit(1, 0).Find(&item); errMsg != nil {
		fmt.Printf(errMsg.Error())
		return
	}

	fmt.Printf("name:%v", item[0].Name)
}

// 查询测试
func QueryTst2(engine *xorm.Engine) {
	data := make([]*Tt, 0)

	if errMsg := engine.Sql("select * from Tt").Find(&data); errMsg != nil {
		fmt.Printf(errMsg.Error())
		return
	}

	for _, item := range data {
		fmt.Printf("\r\nId:%v	name:%v", item.Id, item.Name)
	}
}

// 事务操作
func TransactionTst(engine *xorm.Engine) {
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

}
