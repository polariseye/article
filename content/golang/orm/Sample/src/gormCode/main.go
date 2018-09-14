package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	UserId int    `gorm:"primary_key;column:UserId"`
	Name   string `gorm:"column:Name"`
	Sex    byte   `gorm:"column:Sex"`
}

func (this *UserInfo) BeforCreate(scope *gorm.Scope) {
	// scope.SetColumn("UserId", 0)
	fmt.Println("执行了")
}

// 初始化
func init() {

	// 调整表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return "p_" + tableName
	}
}

func main() {

	db := initDb()

	var data []*UserInfo
	Find(db, "p_user_info", &data)

	for _, item := range data {
		fmt.Println("userid=", item.UserId, "	Name=", item.Name, "	Sex=", item.Sex)
	}
}

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/operation_i_dzz?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 记录日志到标准输出
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// 设置连接池信息（通过sql.Db设置）
	db.DB().SetMaxOpenConns(100)

	// 不使用表名的复数形式。如果为false，结构体名为User，则数据库的表名会被改为users
	db.SingularTable(true)

	return db
}

func createTableTst(db *gorm.DB) {
	userInfo := &UserInfo{}

	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(userInfo)
}

func insertTst(db *gorm.DB) {
	userInfo := &UserInfo{
		UserId: 1,
		Name:   "name1",
		Sex:    1,
	}

	// 执行添加数据操作，但是未返回添加结果信息
	result := db.Create(userInfo)
	if result.Error != nil {
		fmt.Println("insert异常，错误信息:", result.Error.Error())
		return
	}

	if result.RowsAffected <= 0 {
		fmt.Println("insert成功，但未影响任何行")
	} else {
		fmt.Println("insert成功，影响记录数:", result.RowsAffected)
	}
}

func updateTst(db *gorm.DB) {
	userInfo := &UserInfo{
		UserId: 1,
		Name:   "name0",
		Sex:    0,
	}

	// result := db.Model(userInfo).Update("Name") // 更新指定字段
	/*
		result := db.Model(userInfo).Updates(map[string]interface{}{
		"UserId": "1",
		"Name":   "name3"}) // 使用字典更新
	*/

	/*
		result := db.Model(userInfo).Select("Name").Updates(map[string]interface{}{
			"UserId": "1",
			"Name":   "name3"}) // 更新选择字段
	*/
	/*
		result := db.Model(userInfo).Omt("Name").Updates(map[string]interface{}{
			"UserId": "1",
			"Name":   "name3"}) // 更新除了Name的所有字段
	*/
	/*
		result := db.Model(userInfo).
		Where("UserId=?", userInfo.UserId).
		UpdateColumn("Name=?", userInfo.Name) // 完全自定义更新
	*/

	result := db.Save(userInfo) // 更新所有字段
	if result.Error != nil {
		fmt.Println("update异常，错误信息:", result.Error.Error())
		return
	}

	if result.RowsAffected <= 0 {
		fmt.Println("update成功，但未影响任何行")
	} else {
		fmt.Println("update成功，影响记录数:", result.RowsAffected)
	}
}

// 删除测试
func deleteTst(db *gorm.DB) {
	userInfo := &UserInfo{
		UserId: 1,
		Name:   "name0",
		Sex:    0,
	}

	result := db.Delete(userInfo) // 会自动加上主键的值作为删除条件
	// result := db.Table("p_user_info").Where("UserId = ?", userInfo.UserId).Delete() //自定义删除
	if result.Error != nil {
		fmt.Println("Delete异常，错误信息:", result.Error.Error())
		return
	}

	if result.RowsAffected <= 0 {
		fmt.Println("Delete成功，但未影响任何行")
	} else {
		fmt.Println("Delete成功，影响记录数:", result.RowsAffected)
	}
}

// 查询测试1
func queryTst(db *gorm.DB) {
	var data []*UserInfo
	result := db.Where("Name like '%name%'"). //// 指定查询条件
							Select("UserId,Name,Sex").     //// 指定筛选字段
							Order("UserId ASC,Name DESC"). //// 指定排序规则
							Limit(10).                     //// 指定limit部分
							Find(&data)                    //// 查询集合
	if result.Error != nil {
		fmt.Println("select异常，错误信息:", result.Error.Error())
		return
	}

	if result.RowsAffected <= 0 {
		fmt.Println("select成功，但未影响任何行")
	} else {
		fmt.Println("select成功，影响记录数:", result.RowsAffected)
	}

	for _, item := range data {
		fmt.Println("userid=", item.UserId, "	Name=", item.Name, "	Sex=", item.Sex)
	}
}

func Find(db *gorm.DB, tableName string, out interface{}) {
	result := db.Table(tableName).Find(out) //// 查询集合
	if result.Error != nil {
		fmt.Println("select异常，错误信息:", result.Error.Error())
		return
	}

	if result.RowsAffected <= 0 {
		fmt.Println("select成功，但未影响任何行")
	} else {
		fmt.Println("select成功，影响记录数:", result.RowsAffected)
	}
}

// 查询测试2
func queryTst2(db *gorm.DB) {
	var data UserInfo
	result := db.First(&data)
	if result.Error != nil {
		fmt.Println("select异常，错误信息:", result.Error.Error())
		return
	}

	if result.RowsAffected <= 0 {
		fmt.Println("select成功，但未影响任何行")
	} else {
		fmt.Println("select成功，影响记录数:", result.RowsAffected)
	}

	fmt.Println("userid=", data.UserId, "	Name=", data.Name, "	Sex=", data.Sex)
}

// 事务处理
func transactionTst(db *gorm.DB) {
	db = db.Begin()
	userInfo := &UserInfo{
		UserId: 1,
		Name:   "name0",
		Sex:    0,
	}

	// 添加记录
	result := db.Create(userInfo)
	if result.Error != nil {
		result.Rollback()
		return
	}

	// 删除记录
	result = result.Table("p_UserInfo").Where("Name=?", "name1").Delete(nil) // 会自动加上主键的值作为删除条件
	if result.Error != nil {
		result.Rollback()
		return
	}

	// 事务提交
	result.Commit()
}
