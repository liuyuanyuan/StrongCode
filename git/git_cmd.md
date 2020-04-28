# Git常用命令

[TOC]

## 删除Github远程仓库中文件夹

```
1.克隆远程仓库到本地库：
git clone git@github.com:xxx/xxx.git
2.对需要删除的文件、文件夹进行如下操作:
git rm test.txt  (删除文件)git rm -r test    (删除文件夹)
3.提交修改：
git commit -m "Delete some files."
4.将修改提交到远程仓库的xxx分支:
git push origin xxx
```



## 分支管理

查看分支详情

```
git branch -a
```

创建分支

```shell
#以当前代码创建新分支
git branch -b newbranch
#将本地分支提交到远程库
git push origin newbranch
```

删除分支

```
git push --delete origin branchname
```



创建与远程分支对应的本地分支并下载代码到本地
git checkout -b remotebranch origin/remotebranch

提交所有分支到远程库origin
git push -u origin --all


撤销git add .操作
git add . （空格 + 点） 表示当前目录所有文件，不小心就会提交其他文件
git add 如果添加了错误的文件，撤销操作如下：

git status 先看一下add 中的文件 
git reset HEAD 如果后面什么都不跟的话 就是上一次add 里面的全部撤销了 
git reset HEAD XXX/XXX/XXX.java 就是对某个文件进行撤销了


代码回退到某一次commit(不可逆)
查看提交历史，找到要回退到的commit id
git log 
回退代码到制定commit id
git reset --hard commitId
再次查看提交历史，commit id之后的历史不存在了
git log


查看所有提交历史
git log
切换到某次提交
git checkout commitID
修改本次commit的message
git commit --amend
进入编辑状态，修改message内容即可（与vim用法类似）。


#修改远程库
git remote rename origin old-origin
git remote add origin git@host:kernel/xx.git
#从远程库拉提交码
git push -u origin --all
git push -u origin --tags #提交包括tag
#向远程库拉取代码,如需拉取tag，需要加--tag
git pull remotename remotebranch:localbranch  


#git 一次add 多个文件的方法
git add file_1 file_2 file_3


#撤销git commit的内容 
1、找到之前提交的git commit的id 
git log 
找到想要撤销的id 
2、git reset –hard id 
完成撤销,同时将代码恢复到前一commit_id 对应的版本 
3、git reset id 
完成Commit命令的撤销，但是不对代码修改进行撤销，可以直接通过git commit 重新提交对本地代码的修改


#git clone失败的几种解决办法
#缓存区溢出
git config http.postBuffer 524288000
#网络下载速度过缓
git config --global http.lowSpeedLimit 0
git config --global http.lowSpeedTime 999999



**分支重命名**

```shell
#本地分支重命名
git branch -m oldName newName
#上传新命名的分支
git push origin newName
#删除远程旧分支
git push --delete origin oldName
```

