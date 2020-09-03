mongodb问题解决笔记
--------------------------------------------

## centos下，服务掉电重启后，mongodb无法启动

**问题描述**

 `systemctl start mongod`后提示`Job for mongod.service failed because the control process exited with error code. See "systemctl status mongod.service" and "journalctl -xe" for details`

`systemctl status mongod.service`打印如下:
````
mongod.service - MongoDB Database Server
   Loaded: loaded (/usr/lib/systemd/system/mongod.service; enabled; vendor preset: disabled)
   Active: failed (Result: exit-code) since Wed 2020-09-02 10:11:06 CST; 1min 23s ago
     Docs: https://docs.mongodb.org/manual
  Process: 3813 ExecStart=/usr/bin/mongod $OPTIONS (code=exited, status=1/FAILURE)
  Process: 3810 ExecStartPre=/usr/bin/chmod 0755 /var/run/mongodb (code=exited, status=0/SUCCESS)
  Process: 3807 ExecStartPre=/usr/bin/chown mongod:mongod /var/run/mongodb (code=exited, status=0/SUCCESS)
  Process: 3804 ExecStartPre=/usr/bin/mkdir -p /var/run/mongodb (code=exited, status=0/SUCCESS)

Sep 02 10:11:06 localhost.localdomain systemd[1]: Starting MongoDB Database Server...
Sep 02 10:11:06 localhost.localdomain mongod[3813]: about to fork child process, waiting until server is ready for connections.
Sep 02 10:11:06 localhost.localdomain mongod[3813]: forked process: 3815
Sep 02 10:11:06 localhost.localdomain mongod[3813]: ERROR: child process failed, exited with error number 1
Sep 02 10:11:06 localhost.localdomain mongod[3813]: To see additional information in this output, start without the "--fork" option.
Sep 02 10:11:06 localhost.localdomain systemd[1]: mongod.service: control process exited, code=exited status=1
Sep 02 10:11:06 localhost.localdomain systemd[1]: Failed to start MongoDB Database Server.
Sep 02 10:11:06 localhost.localdomain systemd[1]: Unit mongod.service entered failed state.
Sep 02 10:11:06 localhost.localdomain systemd[1]: mongod.service failed.

````

**解决步骤**
1. 删除lock后缀的文件:`rm /data/mongodb/*.lock`
2. 修复数据文件:`sudo mongod -f /etc/mongod.conf --repair`
3. 手动启动mongod服务:`sudo mongod -f /etc/mongod.conf --fork`
