## APUE(Unix环境高级编程)例子Golang版实现

### APUE源码编译
我的实验环境是centos-7.0-x86_64虚拟机，vagrant环境下运行。

**安装GCC**

```
yum install gcc
```

获取apue第三版源码并解压

```
mkdir ~/
wget -c "http://www.apuebook.com/src.3e.tar.gz"
tar zxvf src.3e.tar.gz
cd apue.3e
```

设置环境变量

```
echo 'export C_INCLUDE_PATH=~/apue.3e/include' >> $HOME/.bashrc
echo 'export LIBRARY_PATH=~/apue.3e/lib' >> $HOME/.bashrc
source ~/.bashrc
```

**测试**

```
cd fileio
make //编译依赖环境
gcc -ansi -I../include -Wall -DLINUX -D_GNU_SOURCE  seek.c -o seek  -L../lib -lapue
./seek < /etc/passwd
seek OK 
```

至此，APUE编译环境设置OK。

