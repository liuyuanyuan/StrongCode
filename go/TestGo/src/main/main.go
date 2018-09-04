package main


import (
	"../do" //import do package
	//"log"
	
	"net/http"
	
	"net"
	"strings"
)

func init(){
	do.SavePid()
	
	startSocket()
}

func main() {
	go do.CacheDatas()
	
	router := do.NewRouter()
	do.Info.Println(http.ListenAndServe(":8085", router))
}

func startSocket() {
	server, err := net.Listen("tcp", "192.168.100.170:1208")
	if err != nil {
		do.Info.Println("Fail to start server, %s\n", err)
	} else {
		do.Info.Println("Server Started ...")
		go func() {
			for {
				conn, err := server.Accept()
				if err != nil {
					do.Error.Println("Fail to connect, %s\n", err)
					break
				}
				go socketHandler(conn)
			}
		}()
	}
}

func socketHandler(c net.Conn) {
	do.Info.Println("Enter socketHandler")
	if c == nil {
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		do.Info.Println("inStr: ", inStr)
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			do.Info.Println("Unsupported command: %s\n", inputs[0])
		}
	}
	do.Info.Println("Connection from %v closed. \n", c.RemoteAddr())
}



/*
import (
	"log"
	"fmt"
	"runtime"
)
// init 在main 之前调用
func init() {
	log.SetOutput(os.Stdout)// 将日志输出到标准输出
}
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	go say("0") //开一个新的Goroutines执行
	go say("1") //当前Goroutines执行
	say("2")
}
*/

/*
import (
	//"../do"
	"fmt"
	//"time"
	//"log"
	//"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	//  log.SetOutput(os.Stdout)
}
*/

// main is the entry point for the program.
//func main() {
	//var array [2]int
	//array[1] = 1
	//collection.PassArray(&array)

    //do.GetOSCpu()
    
    
    /*ch := make(chan int, 1)
    fmt.Println("== before ==")
    for i := 0; i < 10 ; i++{
	    select {
	    case ch <- i:
         //do nothing
	    case x := <-ch: 
         fmt.Println(x);
    }
    }
    fmt.Println("== after ==")
    */
    
    /*fmt.Println("Commencing countdown.")
    tick := time.Tick(1 * time.Second)
    for countdown := 10; countdown >0; countdown-- {
	    fmt.Println(countdown)
	    <- tick
    }*/
    
    
    
    /*
    ch := make(chan string, 3)
    
    ch <- "A" 
    ch <- "B"
    ch <- "C" 
    
    //fatal error: all goroutines are asleep - deadlock!
    //fmt.Println("abc")    
    //ch <- "D"
    
    fmt.Println(<- ch)//A
    fmt.Println(<- ch)//B
    ch <- "D"
    fmt.Println("-----")
    fmt.Println(cap(ch))//3
    fmt.Println(len(ch))//2
    */
//}

/*
//单向（无缓冲）通道
func main() {   
    naturals := make(chan int)
    squares := make(chan int)
    
    go counter(naturals)
    go squarer(squares, naturals)
    
    printer(squares)
}
*/
/*
func counter(out chan<- int){
	for x :=0; x<10; x++{
		fmt.Println("counter: ", x)
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int){
	for v := range in{
		fmt.Println("squarer: ", v)
		out <- v * v
	}
	close(out)
}
func printer(in <- chan int){
	for v := range in{
		fmt.Println(v)
	}
}
*/
/*
$ go run  TestGo/src/main/main.go
counter:  0
counter:  1
squarer:  0
squarer:  1
counter:  2
0
1
squarer:  2
4
counter:  3
counter:  4
squarer:  3
squarer:  4
counter:  5
9
16
squarer:  5
25
counter:  6
counter:  7
squarer:  6
squarer:  7
counter:  8
36
49
squarer:  8
64
counter:  9
squarer:  9
81
*/









