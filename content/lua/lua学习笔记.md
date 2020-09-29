lua学习笔记
------------------------
#环境搭建
* vscode中直接安装插件:`lua`与`lua debug`即可
* 编写test.lua:
	````
	print("hello world")
	````
	按F5即可调试运行

# 语法
**数据类型**

数据类型|描述
-|-|
nil|只有值nil属于该类，表示一个无效值（在条件表达式中相当于false）
boolean|包含两个值：false和true。
number|表示双精度类型的实浮点数
string|字符串由一对双引号或单引号来表示
function|由 C 或 Lua 编写的函数
userdata|表示任意存储在变量中的C数据结构
thread|表示执行的独立线路，用于执行协同程序
table|Lua 中的表（table）其实是一个"关联数组"（associative arrays），数组的索引可以是数字、字符串或表类型。在 Lua 里，table 的创建是通过"构造表达式"来完成，最简单构造表达式是{}，用来创建一个空表。

字符串说明
  * 多行字符串使用 [[字符串]]，里面的"也会正常输出
  * 字符串连接使用 `..`,例如:print("你好".."呀")
  * 获取字符串长度:`#"你好"`
  * 在对一个数字字符串上进行算术操作时，Lua 会尝试将这个数字字符串转成一个数字

	示例如下:
	````
	local str="nihaoa"
	print("str长度为:"..#str)
	local multMemo=[[
<html>
<head></head>
<body>
    <span>helloworld</span>
</body>
</html>
]]
	print("多行文本:"..multMemo)
	print("字符串加法:"..("1"+2))

	````

table说明
  * table 的创建是通过"构造表达式"来完成，最简单构造表达式是{}
  * table的key可以是`数字、字符串或表类型`