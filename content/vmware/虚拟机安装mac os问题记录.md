虚拟机安装mac os问题记录
------------------------------------------------------------------------------------------

安装参考文档:[https://www.cnblogs.com/emanlee/p/13344131.html](https://www.cnblogs.com/emanlee/p/13344131.html)

**该虚拟机要求使用 AVX2,但 AVX 不存在。因此该虚拟机无法开启**

参考文档:[https://blog.csdn.net/qq_53079406/article/details/122683340](https://blog.csdn.net/qq_53079406/article/details/122683340)<br>
调整`virtualHW.version` 为`10`即可解决

**客户机操作系统已禁用 CPU，请关闭或重置虚拟机**
参考文档:[https://blog.csdn.net/weixin_46585199/article/details/122147713](https://blog.csdn.net/weixin_46585199/article/details/122147713) <br>
修改xxx.vmx文件。在最后添加
````
smc.version = "0"
cpuid.0.eax = "0000:0000:0000:0000:0000:0000:0000:1011"
cpuid.0.ebx = "0111:0101:0110:1110:0110:0101:0100:0111"
cpuid.0.ecx = "0110:1100:0110:0101:0111:0100:0110:1110"
cpuid.0.edx = "0100:1001:0110:0101:0110:1110:0110:1001"
cpuid.1.eax = "0000:0000:0000:0001:0000:0110:0111:0001"
cpuid.1.ebx = "0000:0010:0000:0001:0000:1000:0000:0000"
cpuid.1.ecx = "1000:0010:1001:1000:0010:0010:0000:0011"
cpuid.1.edx = "0000:0111:1000:1011:1111:1011:1111:1111"
smbios.reflectHost = "TRUE"
hw.model = "MacBookPro14,3"
board-id = "Mac-551B86E5744E2388"
keyboard.vusb.enable = "TRUE"
mouse.vusb.enable = "TRUE"
````

**鼠标键盘无法使用**

参考文档: [https://blog.csdn.net/zhoupian/article/details/122659135](https://blog.csdn.net/zhoupian/article/details/122659135)<br>
把虚拟机USB输入设备更改为USB2.0