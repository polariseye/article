unpack failed问题解决
-------------------------
今天`git push origin develop`时，出现提交失败。具体信息如下：

````
Enumerating objects: 261, done.
Counting objects: 100% (188/188), done.
Delta compression using up to 4 threads.
Compressing objects: 100% (136/136), done.
Writing objects: 100% (137/137), 16.50 KiB | 768.00 KiB/s, done.
Total 137 (delta 114), reused 0 (delta 0)
remote: Resolving deltas: 100% (114/114)
error: remote unpack failed: error Missing commit 58f6224c776a1ddcdbc111daa30adc77a6a7c32b
To http://193.168.5.4:8888/r/prj_001/svr.git
 ! [remote rejected]   develop -> develop (n/a (unpacker error))
error: failed to push some refs to 'http://hello@127.0.0.1:8888/r/prj_001/hello.git'
````

百度后得到解决办法:
原因:本地索引出错
解决方法：

1. git gc
2. git pull --rebase（过程中遇到冲突需要先解决冲突，再git add .）
3. git push

# 参考资料
[git提交代码时出现错误：error : unpack failed : error Missing commit XXX](https://blog.csdn.net/irislulimin/article/details/80760400)