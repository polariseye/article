excel操作
--------------------
综合对比了众多excel库以后，我选择了库[openpyxl](http://openpyxl.readthedocs.io/en/stable/)，需要说明的是，这个只支持xlsx格式的，如果需要支持xls，则需要另选
# openpyxl 使用

## 下载
```
pip install openpyxl
```
如果需要往excel里面写入图片，则需要下载[pillow库](https://pypi.python.org/pypi/Pillow/)
```
pip install pillow
```

## 使用
使用参考[这儿](http://liyangliang.me/posts/2013/02/using-openpyxl-to-read-and-write-xlsx-files/)

**读数据**
```
from openpyxl import load_workbook // 导入操作库
wb = load_workbook("e:/dd.xlsx") // 打开excel文件
sheets = wb.get_sheet_names() // 获取excel的所有工作簿
ws = wb.get_sheet_by_name(sheets[0]) // 获取指定的工作簿

# 读取所有数据
for row in ws.rows:
	for cellItem in row:
		print(cellItem.value)

# 读取指定单元格的值
print(ws.cell('B12').value)
print(ws.cell(row=12, column=2).value)

# 获得行和列的个数
print(sheet.max_row, sheet.max_column)
```

**在字母和数字之间转换**

用于列，方便在字母和数字索引之间转换：
```
import openpyxl
from openpyxl.cell import get_column_letter, column_index_from_string
 
print(get_column_letter(1))
# 'A'
print(get_column_letter(2))
# 'B'
print(get_column_letter(27))
# 'AA'
print(get_column_letter(1000))
# 'ALL'
 
workbook = openpyxl.load_workbook('test.xlsx')
sheet = workbook.get_sheet_by_name('Sheet1')
 
print(get_column_letter(sheet.max_column))
# 'C'
print(column_index_from_string('A'))
# 1
print(column_index_from_string('AA'))
# 27
```
**获得 A2 : C4 单元格区域**
```
import openpyxl
 
workbook = openpyxl.load_workbook('test.xlsx')
sheet = workbook.get_sheet_by_name('Sheet1')
 
print(tuple(sheet['A2':'C4']))
print('---------------------------')
 
for rowOfCell in sheet['A2':'C4']:
    for cell in rowOfCell:
        print(cell.coordinate, cell.value)
```

**写数据**
```
from openpyxl import Workbook
 
# 在内存中创建一个workbook对象，而且会至少创建一个 worksheet
wb = Workbook()
 
ws = wb.get_active_sheet()
print ws.title
ws.title = 'New Title'  # 设置worksheet的标题
 
# 设置单元格的值
ws.cell('D3').value = 4
ws.cell(row=3, column=1).value = 6
 
new_ws = wb.create_sheet(title='new_sheet')
for row in range(1, 100):
    for col in range(1, 10):
        new_ws.cell(row=row, column=col).value = row+col
 
# 最后一定要保存！
wb.save(filename='new_file.xlsx')
```

## 格式设置
**设置行高和列宽**
```
import openpyxl
 
workbook = openpyxl.Workbook()
sheet = workbook.active
sheet['A1'] = 'Hello'
sheet['B2'] = 'World'
 
sheet.row_dimensions[1].height = 70
sheet.column_dimensions['B'].width = 20
 
workbook.save('test1.xlsx')
```

**合并单元格**
```
import openpyxl
workbook = openpyxl.Workbook()
sheet = workbook.active
 
sheet.merge_cells('A1:D3')
sheet['A1'] = '单发独立解等法规和里'
 
workbook.save('test1.xlsx')
```
合并单元格的反操作：
```
sheet.unmerge_cells('A1:D3')
```

**锁定行或列**

第一行浮动：
```
import openpyxl
 
workbook = openpyxl.load_workbook('test.xlsx')
sheet = workbook.active
 
sheet.freeze_panes = 'A2'
 
workbook.save('test1.xlsx')
```
A2’冻结第一行；’B1’冻结第一列；

**图表**
```
import openpyxl
 
workbook = openpyxl.Workbook()
sheet = workbook.active
 
# 创建一些单元格
for i in range(1, 11):
    sheet['A' + str(i)] = i
 
refObj = openpyxl.chart.Reference(sheet, min_col=1, min_row=1, max_col=1, max_row=10)
 
seriesObj = openpyxl.chart.Series(refObj, title='First series')
 
chartObj = openpyxl.chart.BarChart()
chartObj.title = 'My Chart'
chartObj.append(seriesObj)
 
sheet.add_chart(chartObj, 'C5')
 
workbook.save('test1.xlsx')
```

# xls操作
xls操作则以标准库的为准:`xlrd`与`xlwd`

 **读数据**
```
#打开excel文件
data=xlrd.open_workbook('data.xlsx')     
#获取第一张工作表（通过索引的方式）
table=data.sheets()[0] 
#data_list用来存放数据
data_list=[]    
#将table中第一行的数据读取并添加到data_list中
data_list.extend(table.row_values(0))
#打印出第一行的全部数据
for item in data_list:
    print item
```

循环遍历所有行
```
for rownum in xrange(table.nrows):
    print table.row_values(rownum)
```

读取单元格
```
cell_A1=table.row(0)[0].value
#或者像下面这样
cell_A1=table.cell(0,0).value
#或者像下面这样通过列索引
cell_A1=table.col(0)[0].value

```

**写数据**

`xlwd`不能写入超过65535行、256列的数据


# 参考资料
* [文强-Python操作excel的几种方式](http://wenqiang-china.github.io/2016/05/13/python-opetating-excel/)
* [斗大的熊猫-Python操作Excel表格(openpyxl)](http://blog.topspeedsnail.com/archives/5404)