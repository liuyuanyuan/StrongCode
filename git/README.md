# Git 常用命令

[TOC]

#### git clone失败的几种解决办法

```bash
#缓存区溢出
git config http.postBuffer 524288000

#网络下载速度过缓
git config --global http.lowSpeedLimit 0
git config --global http.lowSpeedTime 999999
```



#### remote 管理

```bash
# 重命名remote
git remote rename origin old-origin
# 添加remote
git remote add origin git@host:kernel/xx.git

#2 往远程库拉提交码
git push -u origin --all
git push -u origin --tags #提交包括tag

#向远程库拉取代码，如需拉取tag，需要加--tag
git pull remotename remotebranch:localbranch  
```



#### 删除Github远程仓库中文件夹

```bash
1.克隆远程仓库到本地库：
git clone git@github.com:xxx/xxx.git

2.对需要删除的文件、文件夹进行如下操作:
git rm test.txt  (删除文件)git rm -r test    (删除文件夹)

3.提交修改：
git commit -m "Delete some files."

4.将修改提交到远程仓库的xxx分支:
git push origin xxx
```



#### 分支管理

查看分支详情

```
git branch -a
git branch -v
```

创建分支

```shell
#以当前代码创建新分支
git branch -b newbranch
#将本地分支提交到远程库
git push origin newbranch
```

删除分支

```bash
git push --delete origin branchname
```

创建与远程分支对应的本地分支并下载代码到本地

```bash
git checkout -b remotebranch origin/remotebranch
```

提交所有分支到远程库origin

```bash
git push -u origin --all
```

分支重命名

```shell
#本地分支重命名
git branch -m oldName newName
#上传新命名的分支
git push origin newName
#删除远程旧分支
git push --delete origin oldName
```



#### 撤销git add .操作

```bash
git add .  #（空格 + 点） 表示当前目录所有文件，不小心就会提交其他文件
git add    # 如果添加了错误的文件，撤销操作如下：
git add  file1 file2 file3  # 一次添加多个文件

git status      # 先看一下add 中的文件 
git reset HEAD  # 如果后面什么都不跟的话 就是上一次add 里面的全部撤销了 
git reset HEAD XXX/XXX/XXX.java # 就是对某个文件进行撤销了
```



#### 代码回退到某一次commit(不可逆)

```bash
#查看提交历史，找到要回退到的commit id
git log 

#回退代码到制定commit id
git reset --hard commitId

#强行将本地回退推到远程库
git push origin xx_branch --force

#再次查看提交历史，commit id之后的历史不存在了
git log
```



#### 修改已提交commit的message

```bash
# 查看所有提交历史
git log
# 切换到某次提交
git checkout commitID
# 修改本次commit的message
git commit --amend
# 进入编辑状态，修改message内容即可（与vim用法类似）。
```

