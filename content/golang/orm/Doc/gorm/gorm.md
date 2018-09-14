[gorm](https://jasperxu.github.io/gorm-zh/)

----------------------------------------------------------
# 简介
项目地址：[点这里](https://github.com/jinzhu/gorm)

* 全功能ORM（几乎）
* 关联（包含一个，包含多个，属于，多对多，多种包含）
* Callbacks（创建/保存/更新/删除/查找之前/之后）
* 预加载（急加载）
* 事务
* 复合主键
* SQL Builder
* 自动迁移
* 日志
* 可扩展，编写基于GORM回调的插件
* 每个功能都有测试
* 开发人员友好

## 模型定义
	type UserInfo struct {
		UserId int    `gorm:"primary_key;column:UserId"`
		Name   string `gorm:"column:Name"`
		Sex    byte   `gorm:"column:Sex"`
	}

说明:

1. 使用primary_key指定int为主键会默认认为是自增主键
2. 字段映射采用的是蛇形映射规则，UserId会被映射为user_id，如果不强制指定，则会映射到不存在的字段

## 初始化

首先需要全局初始化，用于注册表名映射函数，比如添加表名前缀后缀

	// 初始化
	func init() {
	
		// 调整表名
		gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
			return "p_" + tableName
		}
	}

然后初始化数据库连接信息

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

说明：

1. 由于是注册的是全局表名调整，所以无法分别为每个连接设置表名调整函数
2. 默认表名映射为struct的复数形式，这个太强制了，应该首先考虑的是常规映射

## 查询

	var data []*UserInfo
	result := db.Where("Name like '%name%'").  //// 指定查询条件
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

说明：

1. 使用Find查询集合，First、Last查询指定的一条数据
2. 字段名映射存在大小写区分，大小写不匹配会出错

## 插入
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

## 更新
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

## 删除
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

## 总结
**优点**

1. 对表结构的调整具有良好的支持
2. 每次基于gorm.DB的调用均是返回的新的gorm.Scope，这使得每次接口调用均是协程安全的

**存在问题**

1. 数据库字段名是按照蛇形命名，暂时发现只能通过tag标签调整，而且要求大小写完全匹配
2. 如果把数值字段设置为主键，会主动加上自增，暂未找到办法去除
3. 表名调整是全局的，无法迎合现在gameserver对表名调整的需求