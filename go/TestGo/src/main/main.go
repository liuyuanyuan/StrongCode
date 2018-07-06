package main


import (
	"../do" //import do package
	"log"
	"net/http"
)

func main() {
	router := do.NewRouter()
	log.Fatal(http.ListenAndServe(":8085", router))
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
	"../do"
	"fmt"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	//		var array [2]int
	//		array[1] = 1
	//		collection.PassArray(&array)

	fmt.Println(do.QDatSession())
}
*/