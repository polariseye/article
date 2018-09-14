日期和时间处理
-----------------------
日期核时间，python有`time`和`calendar`,时间单位以秒为单位

**得到当前时间**

```
val = time.time() # 得到的是时间戳

localtime = time.localtime(time.time()) # 转换成时间元祖的时间

输出结果:
time.struct_time(tm_year=2018, tm_mon=4, tm_mday=7, tm_hour=18, tm_min=51, tm_sec=35, tm_wday=5, tm_yday=97, tm_isdst=0)

```

**时间格式化为字符串**
``` 
time.asctime( time.localtime(time.time()) )
输出结果:
'Sat Apr  7 18:54:29 2018'
```

按照自定义分方式格式化时间
```
time.strftime(format[, t])
```

时间格式化的参数:

* %y 两位数的年份表示（00-99）
* %Y 四位数的年份表示（000-9999）
* %m 月份（01-12）
* %d 月内中的一天（0-31）
* %H 24小时制小时数（0-23）
* %I 12小时制小时数（01-12）
* %M 分钟数（00=59）
* %S 秒（00-59）
* %a 本地简化星期名称
* %A 本地完整星期名称
* %b 本地简化的月份名称
* %B 本地完整的月份名称
* %c 本地相应的日期表示和时间表示
* %j 年内的一天（001-366）
* %p 本地A.M.或P.M.的等价符
* %U 一年中的星期数（00-53）星期天为星期的开始
* %w 星期（0-6），星期天为星期的开始
* %W 一年中的星期数（00-53）星期一为星期的开始
* %x 本地相应的日期表示
* %X 本地相应的时间表示
* %Z 当前时区的名称
* %% %号本身

例如:
```
# 格式化成2016-03-20 11:45:39形式
print time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) 

# 格式化成Sat Mar 28 22:24:24 2016形式
print time.strftime("%a %b %d %H:%M:%S %Y", time.localtime()) 
  
# 将格式字符串转换为时间戳
a = "Sat Mar 28 22:24:24 2016"
print time.mktime(time.strptime(a,"%a %b %d %H:%M:%S %Y")) 
```