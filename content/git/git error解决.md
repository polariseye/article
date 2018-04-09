git出错手动恢复记录
----------------
又是一天加班时，高高兴兴地去吃了一碗面，回来happy了，git出错了
```
$ git branch
fatal: Failed to resolve HEAD as a valid ref.
```

度娘寻他千百度，看了文章[git操作错误Failed to resolve HEAD as a valid ref解决方案](https://blog.csdn.net/ii_bat/article/details/61192230)后，发现现象不一样，但最终还是产生了些灵感。

首先使用fsck看看到底什么情况了
```bash
$ git fsck
error: Invalid HEAD
Checking object directories: 100% (256/256), done.
Checking objects: 100% (6616/6616), done.
error: refs/heads/develop_tmp: invalid sha1 pointer 0000000000000000000000000000000000000000
error: bad ref for .git/logs/HEAD
error: bad ref for .git/logs/refs/heads/develop_tmp
error: bad signature
fatal: index file corrupt
```

这个大概意思是出了很多错：

1. sha1值不正确，变成了0000000000000000000000000000000000000000
2. log日志解析有问题
3. index文件解析有问题

既然问题大概知道了，那咱就手动挨个恢复(我的代码分支名是:develop_tmp)

1. 打开文件./git/refs/heads/develop_tmp 发现是乱码~~~，说明这儿有问题的，正常情况下应该是一个sha1的字符串.先把其他分支的sha1值填充到这个文件
2. 现在咱git log看看，哇哦！有东西了，但发现最近的提交信息不见了， 但git reflog 还有保存最近的提交日志。于是咱就checkout吧，此时错误信息如下: 
```bash
$ git checkout  HEAD@{0}
warning: Log for ref HEAD unexpectedly ended on Mon, 9 Apr 2018 18:31:49 +0800.
error: bad signature
fatal: index file corrupt
```
大意也就是说索引文件出错了，打开文件 *.git/index* ，发现内容如下(全是0值的字节)
```
NULNULNULNULNULNULNULNULNULNULNULNULNULNULNULNUL......
```
3. ok这个没啥，根据参考的两篇文章大概知道，git是可以自己恢复索引的，于是把文件*.git/index* 删掉，然后再执行 *git reset* ，ok了，所有指令都能正常了
4. 趁着高兴劲儿，我又挨个查看了.git目录的所有文件，发现文件 *.git/logs/refs/heads/develop_tmp* 的最后面仍然是可恶的
```
NULNULNULNULNULNULNULNULNULNULNULNULNULNULNULNUL......
```
到此位置整个git就愉快地正常了。

# 总结
1. git目录的内容挺有意思的，除了index文件外，其他的都是文本文件，所以出了错，其实可以手动更改这些文本文件以恢复正常
2. .git\logs\refs\heads目录:存储了各个分支的提交日志信息
3. .git\refs\heads目录：存储了各个分支当前的版本号ID
4. .git\HEAD文件:存储了当前在编辑的分支名
5. 分支检查指令: git fsck
6. 重建索引指令：git reset

# 参考资料

* [git操作错误Failed to resolve HEAD as a valid ref解决方案](https://blog.csdn.net/ii_bat/article/details/61192230)
* [how to fix GIT error: object file is empty?
](https://stackoverflow.com/questions/11706215/how-to-fix-git-error-object-file-is-empty)