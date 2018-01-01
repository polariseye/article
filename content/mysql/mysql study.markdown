mysql ��Ҫ���ݽ̳�
---------------------------
# ���ݱ��޸�
����ʾ���ı�ṹ��
```
CREATE TABLE `student_info` (         
  `Id` int(11) NOT NULL,              
  `Name` varchar(24) DEFAULT NULL,    
  `Sex` int(11) DEFAULT NULL,         
   PRIMARY KEY (`Id`)                  
) ENGINE=InnoDB DEFAULT CHARSET=utf8;  
```
* ���ı�����student_info ����Ϊstudent_info2
```
ALTER table student_info RENAME student_info2;
```

* �޸����ݿ��ֶ�����:Sex����Ϊbigint
```
alter table `student_info` modify `Sex` bigint null after `Name`;
```

* �ֶ���������Sex����ΪSex1,�������͸���Ϊbigint
```
ALTER table student_info change Sex Sex1 bigint null AFTER `Name`;
```

* ����ֶ�
```
--�﷨
ALTER TABLE tablename ADD [COLUMN] column_definition [FIRST | AFTER col_name]
--���Ĭ��ֵ�ֶ�
ALTER TABLE User ADD Age INT NOT NULL DEFAULT 0;
---����ֵ
auto_increment
```

* ɾ���ֶ�
```
ALTER TABLE tablename DROP [COLUMN] col_name
```

* ɾ����
```
DROP TABLE tablename
```

# �������޸�
��ṹ��
```
CREATE TABLE `product` (
   `proID` int(11) NOT NULL AUTO_INCREMENT COMMENT '��Ʒ������',
   `price` decimal(10,2) NOT NULL COMMENT '��Ʒ�۸�',
   `type` int(11) NOT NULL COMMENT '��Ʒ���(0����,1ʳƷ,2����)',
   `dtime` datetime NOT NULL COMMENT '����ʱ��',
   PRIMARY KEY (`proID`)
 ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='��Ʒ��';

CREATE TABLE `producttype` (
   `ID` int(11) NOT NULL COMMENT '��Ʒ���(0����,1ʳƷ,2����)',
   `amount` int(11)  COMMENT 'ÿ�������Ʒ�ܽ��',
   UNIQUE KEY (`ID`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='��Ʒ����ʽ���ܱ�'
```

* ���ݲ���
```
INSERT INTO product(price,type,dtime) VALUES(10.00,0,now()),(10.00,1,now()),(10.00,1,now()),(20.00,2,now()),(30.00,3,now());
```

* �������ݸ���
```
UPDATE product
SET price='20.00',type=0 
WHERE proID=2;
```

* ��������
```
UPDATE producttype,product
SET producttype.amount=product.price
where product.TYPE = producttype.ID AND product.TYPE=1;
``` 

# �洢���̺ͺ���
��mysql���棬�����ʹ洢���̵�����Ϊ��

1. һ����˵���洢����ʵ�ֵĹ���Ҫ����һ�㣬��������ʵ�ֵĹ�������ԱȽ�ǿ���洢���̣�����ǿ�󣬿���ִ�а����޸ı��һϵ�����ݿ�������û����庯����������ִ��һ���޸�ȫ�����ݿ�״̬�Ĳ���

2. ���ڴ洢������˵���Է��ز��������¼����������ֻ�ܷ���ֵ���߱���󡣺���ֻ�ܷ���һ�����������洢���̿��Է��ض�����洢���̵Ĳ���������IN,OUT,INOUT�������ͣ�������ֻ����IN��~~�洢��������ʱ����Ҫ�������ͣ�����������ʱ��Ҫ�����������ͣ��Һ������б������һ����Ч��RETURN���

3. �洢���̣�����ʹ�÷�ȷ�����������������û����庯�����������÷�ȷ������

4. �洢����һ������Ϊһ�������Ĳ�����ִ�У� EXECUTE ���ִ�У���������������Ϊ��ѯ����һ�����������ã�SELECT���ã������ں������Է���һ�����������������ڲ�ѯ�����λ��FROM�ؼ��ֵĺ��档 SQL����в����ô洢���̣�������ʹ�ú���

## �����洢����
�ٷ��﷨���£�
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
˵����

1. **LANGUAGE SQL**������˵����䲿����SQL��䣬δ�����ܻ�֧���������͵����
2. **[NOT] DETERMINISTIC**�����������߳����Ƕ�ͬ���������������ͬ���Ľ��������Ϊ���ǡ�ȷ���ġ���������ǡ���ȷ�����ġ������û�и���DETERMINISTICҲû�и���NOT DETERMINISTIC��Ĭ�ϵľ���NOT DETERMINISTIC����ȷ���ģ�
3. **CONTAINS SQL**����ʾ�ӳ��򲻰�������д���ݵ���䡣
4. **NO SQL**����ʾ�ӳ��򲻰���SQL���
5. **READS SQL DATA**����ʾ�ӳ�����������ݵ���䣬��������д���ݵ����
6. **MODIFIES SQL DATA**����ʾ�ӳ������д���ݵ����
7. **SQL SECURITY DEFINER**����ʾִ�д洢�����еĳ������ɴ����ô洢���̵��û���Ȩ����ִ��
8. **SQL SECURITY INVOKER**����ʾִ�д洢�����еĳ������ɵ��øô洢���̵��û���Ȩ����ִ�С�����������Ĵ洢������д�����ɵ��øô洢���̵��û���Ȩ����ִ�У���ǰ�洢������������ѯEmployee������ҵ�ǰִ�д洢���̵��û�û�в�ѯEmployee���Ȩ����ô�ͻ᷵��Ȩ�޲���Ĵ����������DEFINER����洢��������ROOT�û�������ô�κ�һ���û�������ô洢���̶�����ִ�У���Ϊִ�д洢���̵�Ȩ�ޱ����root��
9. **COMMENT 'string'**:��ע���ʹ�������ֶα�עһ��

ʾ������
```
#�������ݿ�
DROP DATABASE IF EXISTS Dpro;
CREATE  DATABASE Dpro
CHARACTER SET utf8
;

USE Dpro;

#�������ű�
DROP TABLE IF EXISTS Employee;
CREATE TABLE Employee
(id INT NOT NULL PRIMARY KEY COMMENT '����',
 name VARCHAR(20) NOT NULL COMMENT '����',
 depid INT NOT NULL COMMENT '����id'
);

#�����������
INSERT INTO Employee(id,name,depid) VALUES(1,'��',100),(2,'��',101),(3,'��',101),(4,'��',102),(5,'��',103);

#�����洢����
DROP PROCEDURE IF EXISTS Pro_Employee;
DELIMITER $$
CREATE PROCEDURE Pro_Employee(IN pdepid VARCHAR(20),OUT pcount INT )
READS SQL DATA
SQL SECURITY INVOKER
BEGIN
SELECT COUNT(id) INTO pcount FROM Employee WHERE depid=pdepid;

END$$
DELIMITER ;

#ִ�д洢����
CALL Pro_Employee(101,@pcount);

SELECT @pcount;
```

˵����

1. �ڴ����洢���̵�ʱ��һ�㶼����**DELIMITER$$.....END$$ DELIMITER ;**���ڿ�ͷ�ͽ�����Ŀ�ľ��Ǳ���mysql�Ѵ洢�����ڲ���";"���ͳɽ������ţ����ͨ����DELIMITER ;������֪�洢���̽�����

# ����
## declare�������
```
DECLARE var_name[,...] type [DEFAULT value]
```
�ڴ洢���̺ͺ�����ͨ��declare���������BEGIN...END�У��������֮ǰ�����ҿ���ͨ���ظ�������������**declare����ı��������ܴ���@������**

## set�����ֵ����
SET���˿��Ը��Ѿ�����õı�����ֵ�⣬������ָ����ֵ�������±�������SET����ı��������Դ���@�����ţ�SET����λ��Ҳ����BEGIN ....END֮������֮ǰ

��ֵ����:
```
SET var_name = expr [, var_name = expr] ...
```
���岢��ֵ����
```
SET @var_name = expr [, @var_name = expr] ...
```
## SELECT ... INTO��丳ֵ����
 ͨ��select into�����Խ�ֵ���������Ҳ����֮�佫��ֵ��ֵ�洢���̵�out����,����Ĵ洢����select into����֮�佫ֵ����out������

## ����
����������һ�����ڶ�ָ�������Ĵ��������������������ظ��������������� 
### ��������
 ���������������ȶ���ĳ�ִ���״̬����sql״̬�����ƣ�Ȼ��Ϳ������ø��������ƿ�������������������һ���õıȽ��٣�һ���ֱ�ӷ��������������档

```
DECLARE condition_name CONDITION FOR condition_value
 
condition_value:
    SQLSTATE [VALUE] sqlstate_value
  | mysql_error_code
```

# ��������
1. [Mysql�洢���̺ͺ����������](http://www.jb51.net/article/48317.htm)
2. [MySQL �洢���̺ͺ���](http://www.cnblogs.com/chenmh/p/5201473.html)
3. [MySQL ���õ�UPDATE����](http://www.cnblogs.com/chenmh/p/5013606.html)