# SELinux

但在CentOS 5.x之后，SELinux已经是一个很完备的内核模块了，因为此后的SELinux的命令与机制更完善，更易操作管理。但是，并非所有的Linux distributions都支持SELinux，使用之前要先确认好。基于的SELinux的作用，建议不要关闭系统的SELinux。

本章主要介绍：

- SELinux是什么？
- SELinux的运行模式
- SELinux的查看与启动、关闭
- SELinux的必备服务setroubleshoot
- SELinux的策略与规则管理

## SELinux是什么？

SELinux是**Security Enhanced Linux**的缩写，字面意为**安全强化的Linux**。

##### SELinux的设计目标：避免资源的误用。

SELinux是由美国国家安全局（NAS)开发的。因为很多企业发现，通常系统出问题的原因大部分是“内部员工的资源误用”所导致的，实际由外部发动的攻击反而没那么严重。

现在已知所有的系统资源都是通过程序来访问的，那么，如果/var/www/html/设置为777，则代表所有程序均可对该目录访问，万一启动WWW服务软件，那么该软件所有触发的进程将可以写入该目录，而该进程却是对整个Internet提供服务的。如果有通过该进程写入恶意内容，那可真不得了！

由于Linux是自由软件，程序代码都是公开的，因此NAS便使用它作研究目标，最后将研究成果整合到Linux内核中，那就是SELinux。所以说，**SELinux是整合到Linux内核中的一个模块。**更多详情可以在NSA官网搜索selinux：https://www.nsa.gov/What-We-Do/Research/SELinux/。

也就是说，**其实SELinux是在进行程序、文件等权限设置依据的一个内核模块。由于启动网络服务等也是程序，因此刚好也是能够控制网络服务能否访问系统资源的一道关卡。**

##### 传统的文件权限与账号关系：自主访问控制DAC

依据进程的**所有者和资源的rwx权限**，来决定有无访问能力，这样访问文件系统的方式称为**自主访问控制（Discretionary Access Control, DAC）**。

**root具有最高权限**，如果某个程序的不小心被有心人士取得，且该进程属于root权限，那么这个程序就可以在系统上进行任何资源的访问；如果不小心将某个目录的权限设置为777，由于对任何人的权限编程rwx，因此该目录就会被任何人随意访问。这些问题很严重。

##### 以策略规则指定特定程序读取特定文件：强制访问控制MAC

为了避免DAC容易发生的资源误用问题，因此SELinux导入了**强制访问控制（Mandatory AccessControl， MAC），用于针对特定的进程与特定的文件资源来进行权限的控制**。



## SELinux的运行模式

SELinux是通过MAC的方式来控管进程，它控制的主体是进程，而目标则是该进程能否读取的“文件资源”。

##### 主体（subject）：进程；

##### 目标（Object）：文件系统；

##### 策略（Policy）：

   centos 5.x中仅提供两个主要策略：

- target：针对网络服务限制较多，针对本机机制限制较少（默认策略）；

- strict：完整的SELinx限制，限制方面较为严格。

   建议使用默认策略即可。

##### 安全上下文（Security context）：

安全上下文主要用冒号分隔的三个字段：

```
identity:role:type
身份识别:角色:类型
```

- 身份标识（Identity）：

  | 类型     | 含义                                   |
  | -------- | -------------------------------------- |
  | root     | 表示root的账号身份；                   |
  | system_u | 表示系统程序方面的标识，通常就是进程； |
  | user_u   | 表示一般用户账号相关的身份。           |

- 角色（Role）：该字段用来确定这个数据是属于程序、文件资源还是代表用户。

  | 角色     | 含义                                         |
  | -------- | -------------------------------------------- |
  | object_r | 代笔的是文件或目录等文件资源，这是最常见的。 |
  | system_r | 代表的是进程，一般用户也会被指定为它。       |

- 类型（Type，最重要）：在默认的targeted策略中，Idenify与Role字段基本上是不重要的。重要的是Type。Type在文件与进程的定义不同：

  - 在文件资源（Object）上称为类型（Type）；

  - 在主体程序（Subject）上称为域（domain）。

  domain需要与type搭配，则该程序才能够顺利读取文件资源。



## SELinux的查看与启动、关闭

##### 目前SELinux支持三种模式：

- enforcing：强制模式，代表SELinux正在运行中，且已经开始限制domain/type了；

- permissive：宽容模式，代表SELinux正在运行中，不过仅会有警告信息并不会实际限制domain/type的访问。这种模式可以用作SELinux的调试之用。

- disabled：关闭，SELinux没有实际运行。

##### SELinux的查看和修改

SELinux配置文件/etc/selinux/config：

```
vim /etc/selinux/config
SELINUX=enforcing     # enforcing｜disbaled｜permissive
SELINUXTYPE=targeted  # targeted｜strict
```

查看SELinux当前模式的命令： 

```
getenforce
返回 Enforcing｜Permissive
```

查看SELinux的策略（ploicy）： 

```
sestatus [-vb]

```

##### SELinux的启动和关闭

如果修改策略，则需要重启后才生效；

如果修改模式，由enforcing/permissive改成disabled，或由disabled改成enforcing/permissive，则需要重启后生效；

当从disbaled模式转到启动SELinux模式时，由于系统必须要针对文件写入安全上下文（有时也称作security label）的信息，因此开机过程会多消耗这个时间，并且写完后还要再重启一次。

当在enforcing模式下，遇到因为设置问题导致SELinux让某些服务无法正常运行，此时可以将Enforcing模式改为permissive的模式，让SELinux只会警告无法顺利连接的信息，而不是直接抵挡主体进程的读取权限。

在SELinux运行时，可以通过setenforce切换Enforcing和Permissive两个模式；在SELinux停止时（即disabled模式时），无法使用setenforce切换任何模式。

```
setenforce [0|1]
参数：
0：转成Permissive宽容模式
1：转成Enforcing强制模式
```



## 安全上下文的查看、修改

##### 查看安全上下文设置情况

-Z选项用于查看安全上下文。

```
查看文件的安全上下文：
ls -l -Z

查看进程httpd的安全上下文：
ps aux -Z | grep httpd
```

##### 修改chcon和恢复restorecon安全上下文

chcon通过直接指定的方式来处理安全上下文的类型数据。

```
chcon [-R] [-t type] [-u user] [-r role] 文件或目录
chcon [-R] --reference=范例文件 文件或目录
参数：
-R:
-t:
-u:
-r:
 --reference=范例文件：
```

restorecon使用默认的安全上下文来还原。

```
restorecon [-Rv] 文件或目录
参数：
-R:连同子目录一起修改；
-v:将过程显示到屏幕上。
```



## SELinux的必备服务setroubleshoot

启动SELinux，必备的服务至少要启动setroubleshoot。

**setroubleshoot：将SELinux错误信息写入/var/log/meassages中。**

- 设置开机启动setroubleshoot的方法：

```
#查看setroubleshoot的开机启动设置（默认都是开启的，除非3:off或5:off）
chkconfig --list setroubleshoot
>setroubleshoot 0:off 1:off 3:on 4:on 5:on 6:off
#设置开机时启动setroubleshoot
chkconfig setroubleshoot on
>on 开机时启动； off开机时不启动
```

- 查看/var/log/meassages中的setroubleshoot错误信息：

```
cat /var/log/meassages | grep setroubleshoot
```



## SELinux的策略与规则管理

centos 5.x默认使用targerted策略，seinfo命令可以查询策略提供的相关规则。

```shell
1 seinfo命令可以查询策略提供的相关规则
seinfo [-Atrub]
#列出SELinux在当前策略下的统计状态
seinfo
#列出httpd相关的规则
seinfo -b ｜ grep httpd

2 sesearch查询详细的规则
sesearch -a [-s 主体类型] [-t 目标类型] [-b 布尔值]
#查询目标类型httpd_sys_content_t的相关信息
sesearch -a -t httpd_sys_content_t

3 getsebool 与 setsebool 对SELinux布尔条款查询和修改
getsebool [-a] [布尔值条款]
参数：-a： 列出目前系统中所有布尔值条款设置情况。
setsebool [-p] 布尔值条款=[0|1]
参数：-p：直接将设置值写入配置文件，该设置数据将来生效。

getsebool httpd_enable_homedirs
setsebool -P httpd_enable_homedirs=0

4 默认目录的安全上下文查询与修改
#查看默认目录的安全上下文
semanage {login|user|port|interface|fcontext|translation} -l
semanage fcontext -{a|d|m} {-frst} file_spec

#重置默认目录的安全上下文
restorecon [-Rv] 文件或目录
```

(待续...)