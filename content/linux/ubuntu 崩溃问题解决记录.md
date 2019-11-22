
问题描述
````
/sbin/init no such file or directory
...
/bin/sh:can't access tty;job control turned off
````
1. Boot your linux distribution on USB or CD
2. Live boot your system on the USB or CD key
3. In command line type: sudo mount /dev/sda1 or your dist partition /mnt
4. Run command: sudo chroot /mnt
5. Then run: sudo nano /etc/resolv.conf
6. nameserver 8.8.8.8
7. sudo apt-get install init

说明:
1. 如果发现无法修改`/etc/resolv.conf`，则把此文件更名，然后重新创建一个此名的文件即可
2. 如果执行`sudo apt-get install init`报错:`errors were encountered while processing plymouth`,则

参考资料
* [Kernel panic on boot: run-init: /sbin/init: No such file or directory
](https://askubuntu.com/questions/833735/kernel-panic-on-boot-run-init-sbin-init-no-such-file-or-directory)
* [Errors were encountered while processing 解决方法](https://blog.csdn.net/qingfengxiaosong/article/details/87889995)