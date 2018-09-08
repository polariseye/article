文件与文件夹操作
--------------------------------
常规的文件操作调用`os`模块或`os.path`模块即可，`shutil`是os模块的补充，提供了很多扩展操作，比如文件复制。

# 路径操作
```
# 查看当前目录的绝对路径:
>>> os.path.abspath('.')
'/Users/michael'
# 在某个目录下创建一个新目录，首先把新目录的完整路径表示出来:
>>> os.path.join('/Users/michael', 'testdir')
'/Users/michael/testdir'
# 然后创建一个目录:
>>> os.mkdir('/Users/michael/testdir')
# 删掉一个目录:
>>> os.rmdir('/Users/michael/testdir')
```

路径拆分:

`os.path.split()`函数可以把一个路径拆分为两部分，后一部分总是最后级别的目录或文件名
```
>>> os.path.split('/Users/michael/testdir/file.txt')
('/Users/michael/testdir', 'file.txt')
```

`os.path.splitext()`可以直接让你得到文件扩展名，很多时候非常方便：
```
>>> os.path.splitext('/path/to/file.txt')
('/path/to/file', '.txt')
```

参考资料:
* [廖雪峰python教程](https://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000/001431925324119bac1bc7979664b4fa9843c0e5fcdcf1e000)