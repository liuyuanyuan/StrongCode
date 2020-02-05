# Golang in Using

## Config Golang on Win:

1.install go
download, unzip, add E:\Go\bin EVN PATH;

2.create Go workspace(anywhere)
E:\>tree /f GoWorkspace
E:\GOWORKSPACE
├─bin
│      hello.exe
├─pkg
│  └─windows_amd64
│      ├─github.com
│      │  ├─golang
│      │  │      dep.a
│      │  │
│      │  ├─lib
│      │  │      pq.a
│      │  │
│      │  ├─Unknwon
│      │  │      com.a
│      │  │      log.a
│      │  │
│      │  └─urfave
│      │          cli.a
│      │
│      └─gopkg.in
│          └─fsnotify
│                  fsnotify.v1.a
│
└─src
    ├─github.com
    │  │  guru.exe

3. create and run hello world
In e:\GoWorkspace\src\ create hello folder and hello.go:
#hello.go
package main
import "fmt"
func main() {
    fmt.Printf("hello, world\n")
}

E:\GoWorkspace\src>tree /f hello
  E:\GOWORKSPACE\SRC\HELLO
    hello.go

E:\GoWorkspace\src>cd hello
E:\GoWorkspace\src\hello>go run hello.go
hello, world


4.to get dependency project
execute: go get github.com/Unknwon/bra
need: download and install git, and add C:\Program Files\Git\bin to ENV Path
(pay attention to network when download dependency project).



## Config Eclipse dev with Go

ref: https://blog.csdn.net/youbaopipa/article/details/75530665



## Config Golang on Linux

1. install golang
yum install golang
2. get install path
go env
returns GOROOT=/usr/lib/golang
3. create go workspace
makedir /root/Desktop/GoWorkspace
4. config env
vi ~/.bash_profile
export GOROOT=/usr/lib/golang
export GOPATH=/root/Desktop/GoWorkspace
PATH=$PATH:$GOROOT/bin
export PTAH

source ~/.bash_profile

5. code and run
cd /root/Desktop/GoWorkspace

vi hello.go
package main
import "fmt"
func main(){
	fmt.println("hello")
}

go run hello.go 

  hello



## Deploy Golang App to server

ref：https://www.jianshu.com/p/bfaba9b6d46d
go build main.go
then generated main.exe(win) or main(linux)
./main.exe or ./main


==Go Test xxx_test.go==
highgoer@DESKTOP-1QJVH6I MINGW64 /e/GoWorkspace/src/gopl.io/ch9 (master)
$ cd bank2

highgoer@DESKTOP-1QJVH6I MINGW64 /e/GoWorkspace/src/gopl.io/ch9/bank2 (master)
$ go  test
PASS
ok      gopl.io/ch9/bank2       0.052s




