Golang 规范
----------------------------

# 规范 #

## 代码格式化 ##

 统一使用`gofmt`工具进行格式化

## 包依赖管理 ##

建议使用gomodule包管理方式

## 包与文件规范 ##

** 包命名 **

* 包名应该以小写开头，且遵循骆驼峰命名规则。例如：`dbServer`

* 包名需要和对应文件夹名一致

** 包文档 **

每个包都应该有对包的注释说明，建议在一个单独的文件中进行说明，以免分散的包说明给`go doc`工具带来困扰 ，比如：`doc.go`
````
/*
这是一个数据库的操作包
*/
package dbServer;
````

** 文件命名 **

* 与包命名一致,以小写开头，且遵循骆驼峰命名规则,应该是名词短语，且见名知义。例如：`battleMap.go`。

* 文件命名要能体现文件内容的含义

* 单个文件名应该功能单一一点。同一模块的多个子文件可以通过不同后缀区分,但一定需要有以功能模块命名的文件，作为主代码文件。 以`battle`模块为例:`battle.go`、`battle_Config.go`、`battle_Node.go`

## 常量与变量规范 ##

** 常量**

* 以骆驼峰命名规则命名，首字母根据访问权限确定大小写，应该是名词短语，且见名知义

* 一个功能模块的常量应该集中定义，且放在文件最上部 

* 枚举类常量应该还需要额外遵循以下原则
  * 使用统一前缀区分，且放在同一个`const`代码块内(有些IDE会进行特殊展示)，并为const代码块加以说明，例如：
  ````
	// 性别枚举
	const (
		Sex_Man   = 1 // 男
		Sex_Woman = 2 // 女
	)
  ````
  * 枚举类常量建议多使用类型别名来处理

* 普通常量应该遵循以下原则 
  * 建议以`con_`开头，以便能够比较方便找到常量定义，也能够更好地利用代码智能提示功能

  * 建议合理使用`const`块的形式，这可以使代码更整洁

示例:
````
const Con_MaxBattleId = 1 // 最大战斗Id
const Con_MinMapId    = 2 // 最小地图节点Id

// 地图类型
const (
	MapType_Normal = 1 // 普通地图
	MapType_Scret  = 2 // 秘境地图
)
````
 

** 变量 **

* 以骆驼峰命名规则命名，首字母根据访问权限确定大小写，变量命名应该是名词短语，且见名知义

* 函数内的变量，首字母都应该小写

* 多个变量同时定义时，建议使用`var`块进行定义

* slice或数组建议使用List结尾, map使用Data结尾，chancel使用Ch或Chan结尾，struct实例可以使用Obj结尾或者不用后缀,函数型变量应该以Func结尾

* 重要变量应该加上对应注释，以便能清晰明白其使用目的 

示例:

````

func HelloWorld(name string){
	nameList := make([]string ,0 ,4)
	nameSexData := make(map[string]int)
	waitChan := make(chan bool, 4)
}

````

## 函数规范 ##

** 函数声明 **

* 以骆驼峰命名规则命名，首字母根据访问权限确定大小写，如果包外不需要访问请用小写开头的函数

* 名称应该是动宾短语或动词。应该见名知义 

* 一个功能逻辑的多个函数应该放置在一起，主函数或对外的函数应该是在前面（以体现出函数级的逻辑块）

* 原则上建议每个函数都应该有函数功能文档说明
  * 函数名非常直接，参数又比较简单的函数可以不用添加文档说明

  * 函数参数和返回值都应该有对应注释说明

  * 提供对外的函数必须加上对应说明文档  

* 函数参数应该遵循变量命名规范，且以小写打头

* 函数返回值应该如同函数参数一样都有变量名，以便能够外部能比较简单知道返回值含义，比较简单的函数可以不用

** 函数体 **

* 多个变量同时定义时，可以使用var块定义

* 只有此函数使用的常量可以在函数体内直接定义

* 不建议在函数体内直接使用字面值，而用const定义，再使用。如果仍坚持使用，则需要在其附近对这个字面量加以说明,以便代码更容易被理解

* 如果同一个字面量在两个或以上函数中有使用，必须要定义对应的函数外常量,以便代码更容易被理解

* 相对独立的程序块应该用空行进行隔开，而逻辑紧密相关的代码则放在一起，以便逻辑更加清晰

* 函数返回值的顺序应该是:需要的返回数据,exist，error，这样有助于把重点放在需要的数据上。比如:`func getPlayerInfo(key string)(data *PlayerInfo,exist bool,err error)`

* 如果返回值个数超过3个，应该需要考虑是否封装为一个返回的结构体  

* 函数返回语句和其他逻辑之间建议添加一个空行 

示例:
````
// getHeroSkills 获取学员技能信息
// key:玩家Id
// heroIdxList:学员Id列表
// 返回值:
// addBeSkills:技能列表
// err:错误信息
func getHeroSkills(key string, heroIdxList []int32) (addBeSkills []*TAddBeSkillItem, err error) {
	// 加载学员数据
	heroData := LoadHeros(key)
	if heroData == nil {
		err := fmt.Errorf("[getEquipJewelSkills] can not load heroData data, key: %v.", key)
		mylog.Error(err.Error())

		return nil, err
	}

	// 获取学员技能
	addBeSkills = []*TAddBeSkillItem{}
	for _, heroIdx := range heroIdxList {
		if heroIdx <= 0 {
			continue
		}

		// 获取单个学员的技能
		addbeJewelSkills, err := getOneHeroSkill(key, heroData, heroIdx)
		if err != nil {

			return addBeSkills, err
		}

		if len(addbeJewelSkills) > 0 {
			addBeSkills = append(addBeSkills, addbeJewelSkills...)
		}
	}

	return
}

````  

## struct规范 ##

* 以骆驼峰命名规则命名，首字母根据访问权限确定大小写，如果包外不需要访问请用小写开头的函数,GameServer开发时，建议使用T开头(因为每个包内容都太多，如果不加以区分很难找到需要的struct)，如:``type TPlayerInfo struct{}``

* 不要以空泛的名称作为strcut名，否则很难理解,且strcut名应该是名词或名词短语 

* 每个strcut都应该有说明文档

* strcut的成员函数都应该的receiver建议使用`this`，因为这样更容易理解些，但需要注意的是，这个等同于直接把receiver作为函数第一个参数，所以strcut为nil时，仍可以正常调用对应成员函数

 
## interface规范 ##

* 以骆驼峰命名规则命名，首字母根据访问权限确定大小写，如果包外不需要访问请用小写开头的函数

* 建议使用字母`I`开头 

## 关于 error ##

* 建议在error产生的地方记录日志，而不是在最上层记录,因为大多数函数并不知道自己就是最上层，这将导致到处都记录日志，例如
````
func QueryDb(sqlString string)(result []*PlayerInfo,err error){
	result,err = db.Select(sql)
	if err !=nil {
		myLog.Error(err)
	}
}
````

* 非最上层函数，不建议轻易吞掉error，因为，这会导致上层无法知道是出错了，还是什么其他原因，比如:
````
func LoadPlayerInfo(key string)(data *PlayerInfo){
	data=new(PlayerInfo)
	if err:=db.Query(key,data);err!=nil{
		return nil //// 此处会导致上层不知道是没查到数据还是出错了
	}
	
	return
}
````

## 关于slice ##

* 初始化slice建议使用make，而不是new。因为make可以指定内存重新开辟规则,从而可以提升slice效率，如:``var nameList = make([]string ,0 ,4)``


## 其他

* 尽早return：一旦有错误发生，马上返回。以使逻辑更加清晰

举例：

````
不要使用
	if err != nil {
	    // error handling
	} else {
	    // normal code
	}


而推荐使用：
	if err != nil {
	     // error handling
	     return // or continue, etc.
	}
	
	// normal code
````

* 应尽量减少函数中代码的层次，如果层次太多可以使用提前返回方式减少循环

## 参考资料
* [golang编程规范2](https://blog.csdn.net/tanzhe2017/article/details/80924660)