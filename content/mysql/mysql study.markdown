mysql 重要内容整理
---------------------------
# 数据表修改
以下示例的表结构：
```
CREATE TABLE `student_info` (         
  `Id` int(11) NOT NULL,              
  `Name` varchar(24) DEFAULT NULL,    
  `Sex` int(11) DEFAULT NULL,         
   PRIMARY KEY (`Id`)                  
) ENGINE=InnoDB DEFAULT CHARSET=utf8;  
```
* 更改表名：student_info 更改为student_info2
```
ALTER table student_info RENAME student_info2;
```

* 修改数据库字段类型:Sex更改为bigint
```
alter table `student_info` modify `Sex` bigint null after `Name`;
```

* 字段重命名：Sex更名为Sex1,并把类型更改为bigint
```
ALTER table student_info change Sex Sex1 bigint null AFTER `Name`;
```

* 添加字段
```
--语法
ALTER TABLE tablename ADD [COLUMN] column_definition [FIRST | AFTER col_name]
--添加默认值字段
ALTER TABLE User ADD Age INT NOT NULL DEFAULT 0;
---自增值
auto_increment
```

* 删除字段
```
ALTER TABLE tablename DROP [COLUMN] col_name
```

* 删除表
```
DROP TABLE tablename
```

# 表数据修改
表结构：
```
CREATE TABLE `product` (
   `proID` int(11) NOT NULL AUTO_INCREMENT COMMENT '商品表主键',
   `price` decimal(10,2) NOT NULL COMMENT '商品价格',
   `type` int(11) NOT NULL COMMENT '商品类别(0生鲜,1食品,2生活)',
   `dtime` datetime NOT NULL COMMENT '创建时间',
   PRIMARY KEY (`proID`)
 ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品表';

CREATE TABLE `producttype` (
   `ID` int(11) NOT NULL COMMENT '商品类别(0生鲜,1食品,2生活)',
   `amount` int(11)  COMMENT '每种类别商品总金额',
   UNIQUE KEY (`ID`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品类别资金汇总表'
```

* 数据插入
```
INSERT INTO product(price,type,dtime) VALUES(10.00,0,now()),(10.00,1,now()),(10.00,1,now()),(20.00,2,now()),(30.00,3,now());
```

* 单表数据更新
```
UPDATE product
SET price='20.00',type=0 
WHERE proID=2;
```

* 关联更新
```
UPDATE producttype,product
SET producttype.amount=product.price
where product.TYPE = producttype.ID AND product.TYPE=1;
``` 

# 存储过程和函数
在mysql里面，函数和存储过程的区别为：

1. 一般来说，存储过程实现的功能要复杂一点，而函数的实现的功能针对性比较强。存储过程，功能强大，可以执行包括修改表等一系列数据库操作；用户定义函数不能用于执行一组修改全局数据库状态的操作

2. 对于存储过程来说可以返回参数，如记录集，而函数只能返回值或者表对象。函数只能返回一个变量；而存储过程可以返回多个。存储过程的参数可以有IN,OUT,INOUT三种类型，而函数只能有IN类~~存储过程声明时不需要返回类型，而函数声明时需要描述返回类型，且函数体中必须包含一个有效的RETURN语句

3. 存储过程，可以使用非确定函数，不允许在用户定义函数主体中内置非确定函数

4. 存储过程一般是作为一个独立的部分来执行（ EXECUTE 语句执行），而函数可以作为查询语句的一个部分来调用（SELECT调用），由于函数可以返回一个表对象，因此它可以在查询语句中位于FROM关键字的后面。 SQL语句中不可用存储过程，而可以使用函数

## 创建存储过程或函数
官方语法如下：
```
CREATE PROCEDURE sp_name ([proc_parameter[,...]])
    [characteristic ...] routine_body
 
CREATE FUNCTION sp_name ([func_parameter[,...]])
    RETURNS type
    [characteristic ...] routine_body
    
proc_parameter:
[ IN | OUT | INOUT ] param_name type

func_parameter:
param_name type
 
type:
    Any valid MySQL data type
 
characteristic:
    LANGUAGE SQL
  | [NOT] DETERMINISTIC
  | { CONTAINS SQL | NO SQL | READS SQL DATA | MODIFIES SQL DATA }
  | SQL SECURITY { DEFINER | INVOKER }
  | COMMENT 'string'
 
routine_body:
    Valid SQL procedure statement or statements
```
说明：

1. **LANGUAGE SQL**：用来说明语句部分是SQL语句，未来可能会支持其它类型的语句
2. **[NOT] DETERMINISTIC**：如果程序或线程总是对同样的输入参数产生同样的结果，则被认为它是“确定的”，否则就是“非确定”的。如果既没有给定DETERMINISTIC也没有给定NOT DETERMINISTIC，默认的就是NOT DETERMINISTIC（非确定的）
3. **CONTAINS SQL**：表示子程序不包含读或写数据的语句。
4. **NO SQL**：表示子程序不包含SQL语句
5. **READS SQL DATA**：表示子程序包含读数据的语句，但不包含写数据的语句
6. **MODIFIES SQL DATA**：表示子程序包含写数据的语句
7. **SQL SECURITY DEFINER**：表示执行存储过程中的程序是由创建该存储过程的用户的权限来执行
8. **SQL SECURITY INVOKER**：表示执行存储过程中的程序是由调用该存储过程的用户的权限来执行。（例如上面的存储过程我写的是由调用该存储过程的用户的权限来执行，当前存储过程是用来查询Employee表，如果我当前执行存储过程的用户没有查询Employee表的权限那么就会返回权限不足的错误，如果换成DEFINER如果存储过程是由ROOT用户创建那么任何一个用户登入调用存储过程都可以执行，因为执行存储过程的权限变成了root）
9. **COMMENT 'string'**:备注，和创建表的字段备注一样

示例代码
```
#创建数据库
DROP DATABASE IF EXISTS Dpro;
CREATE  DATABASE Dpro
CHARACTER SET utf8
;

USE Dpro;

#创建部门表
DROP TABLE IF EXISTS Employee;
CREATE TABLE Employee
(id INT NOT NULL PRIMARY KEY COMMENT '主键',
 name VARCHAR(20) NOT NULL COMMENT '人名',
 depid INT NOT NULL COMMENT '部门id'
);

#插入测试数据
INSERT INTO Employee(id,name,depid) VALUES(1,'陈',100),(2,'王',101),(3,'张',101),(4,'李',102),(5,'郭',103);

#创建存储过程
DROP PROCEDURE IF EXISTS Pro_Employee;
DELIMITER $$
CREATE PROCEDURE Pro_Employee(IN pdepid VARCHAR(20),OUT pcount INT )
READS SQL DATA
SQL SECURITY INVOKER
BEGIN
SELECT COUNT(id) INTO pcount FROM Employee WHERE depid=pdepid;

END$$
DELIMITER ;

#执行存储过程
CALL Pro_Employee(101,@pcount);

SELECT @pcount;
```

说明：

1. 在创建存储过程的时候一般都会用**DELIMITER$$.....END$$ DELIMITER ;**放在开头和结束，目的就是避免mysql把存储过程内部的";"解释成结束符号，最后通过“DELIMITER ;”来告知存储过程结束。

# 变量
## declare定义变量
```
DECLARE var_name[,...] type [DEFAULT value]
```
在存储过程和函数中通过declare定义变量在BEGIN...END中，且在语句之前。并且可以通过重复定义多个变量。**declare定义的变量名不能带‘@’符号**

## set定义或赋值变量
SET除了可以给已经定义好的变量赋值外，还可以指定赋值并定义新变量，且SET定义的变量名可以带‘@’符号，SET语句的位置也是在BEGIN ....END之间的语句之前

赋值变量:
```
SET var_name = expr [, var_name = expr] ...
```
定义并赋值变量
```
SET @var_name = expr [, @var_name = expr] ...
```
## SELECT ... INTO语句赋值变量
 通过select into语句可以将值赋予变量，也可以之间将该值赋值存储过程的out参数,上面的存储过程select into就是之间将值赋予out参数。

## 变量注意事项
1. 遍历不区分大小些
2. 变量的定义必须写在复合语句的开头，并且在任何其他语句的前面

## 条件
条件的作用一般用在对指定条件的处理，比如我们遇到主键重复报错后该怎样处理。 
### 条件的定义
 定义条件就是事先定义某种错误状态或者sql状态的名称，然后就可以引用该条件名称开做条件处理，定义条件一般用的比较少，一般会直接放在条件处理里面。

```
DECLARE condition_name CONDITION FOR condition_value
 
condition_value:
    SQLSTATE [VALUE] sqlstate_value
  | mysql_error_code
```
###  条件的处理
```
DECLARE handler_type HANDLER FOR condition_value[,...] sp_statement
handler_type:
    CONTINUE
    | EXIT
    | UNDO
condition_value:
    SQLSTATE [VALUE] sqlstate_value
| condition_name
| SQLWARNING
| NOT FOUND
| SQLEXCEPTION
| mysql_error_code
```
处理示例：
```
delimiter $$
CREATE PROCEDURE actor_insert ()
BEGIN
    DECLARE CONTINUE HANDLER FOR SQLSTATE '23000' SET @x2 = 1;
    SET @x = 1;
    INSERT INTO actor(actor_id,first_name,last_name) VALUES(201,'Test','201');
    SET @x = 2;
    INSERT INTO actor(actor_id,first_name,last_name) VALUES (1,'Test','1');
    SET @x = 3;
END;
$$
```

说明:

1. handler_type 现在还只支持 CONTINUE 和 EXIT 两种,CONTINUE 表示继续执行下面的语句，
EXIT 则表示执行终止
2. condition_value 的值可以是通过 DECLARE 定义的 condition_name，可以是 SQLSTATE 的值或
者 mysql-error-code 的值或者 SQLWARNING、NOT FOUND、SQLEXCEPTION，这 3 个值是 3 种
定义好的错误类别，分别代表不同的含义

# 游标
在存储过程和函数中可以使用光标对结果集进行循环的处理。光标的使用包括光标的声明、OPEN、FETCH 和 CLOSE。

声明游标:
```
DECLARE cursor_name CURSOR FOR select_statement
```
打开游标:
```
OPEN cursor_name
```
Fetch游标:
```
FETCH cursor_name INTO var_name [, var_name] ...
```
Close游标:
```
CLOSE cursor_name
```

示例代码:
```
mysql> delimiter $$
mysql>
mysql> CREATE PROCEDURE payment_stat ()
-> BEGIN
-> DECLARE i_staff_id int;
-> DECLARE d_amount decimal(5,2);
-> DECLARE cur_payment cursor for select staff_id,amount from payment;
-> DECLARE EXIT HANDLER FOR NOT FOUND CLOSE cur_payment;
->
-> set @x1 = 0;
-> set @x2 = 0;
->
-> OPEN cur_payment;
->
-> REPEAT
-> FETCH cur_payment INTO i_staff_id, d_amount;
-> if i_staff_id = 2 then
-> set @x1 = @x1 + d_amount;
-> else
-> set @x2 = @x2 + d_amount;
-> end if;
-> UNTIL 0 END REPEAT;
->
-> CLOSE cur_payment;
->
-> END;
-> $$
Query OK, 0 rows affected (0.00 sec)
mysql> delimiter ;
```

注意事项:
* **变量、条件、处理程序、光标都是通过 DECLARE 定义的，它们之间是有先后顺序的要
求的。变量和条件必须在最前面声明，然后才能是光标的声明，最后才可以是处理程序
的声明**

# if语句
```
IF search_condition THEN statement_list
    [ELSEIF search_condition THEN statement_list] ...
    [ELSE statement_list]
END IF
```
# case语句
```
CASE case_value
    WHEN when_value THEN statement_list
    [WHEN when_value THEN statement_list] ...
    [ELSE statement_list]
END CASE
Or:
CASE
    WHEN search_condition THEN statement_list
    [WHEN search_condition THEN statement_list] ...
    [ELSE statement_list]
END CASE
```
# loop语句
```
[begin_label:] LOOP
    statement_list
END LOOP [end_label]
```
这个通常需要Leave语句实现退出循环

# leave语句，用于跳出循环
```
mysql> CREATE PROCEDURE actor_insert ()
-> BEGIN
-> set @x = 0;
-> ins: LOOP
-> set @x = @x + 1;
-> IF @x = 100 then
-> leave ins;
-> END IF;
-> INSERT INTO actor(first_name,last_name) VALUES ('Test','201');
-> END LOOP ins;
-> END;
-> $$
```
# ITERATE 语句，相当于C语言的continue
ITERATE 语句必须用在循环中，作用是跳过当前循环的剩下的语句，直接进入下一轮循环
```
mysql> CREATE PROCEDURE actor_insert ()
-> BEGIN
-> set @x = 0;
-> ins: LOOP
-> set @x = @x + 1;
-> IF @x = 10 then
-> leave ins;
-> ELSEIF mod(@x,2) = 0 then
-> ITERATE ins;
-> END IF;
-> INSERT INTO actor(actor_id,first_name,last_name) VALUES (@x+200,'Test',@x);
-> END LOOP ins;
-> END;
-> $$
```
# REPEAT 语句,相当于C语言的do while
有条件的循环控制语句，当满足条件的时候退出循环
```
[begin_label:] REPEAT
    statement_list
UNTIL search_condition
END REPEAT [end_label]
```
#  WHILE 语句
```
[begin_label:] WHILE search_condition DO
    statement_list
END WHILE [end_label]
```

# 引用资料
1. [Mysql存储过程和函数区别介绍](http://www.jb51.net/article/48317.htm)
2. [MySQL 存储过程和函数](http://www.cnblogs.com/chenmh/p/5201473.html)
3. [MySQL 常用的UPDATE操作](http://www.cnblogs.com/chenmh/p/5013606.html)
4. [深入浅出Mysql]()