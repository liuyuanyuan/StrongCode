package do

import (
	"io"
//	"io/ioutil"
	"log"
	"os"
)

/**
Ref 《Go语言实战》-定制日志记录器
*/
var (
	//Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warn    *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

func init() {
	file, err := os.OpenFile("agent_error.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	//Trace = log.New(ioutil.Discard,
	//	"TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout,
		"INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout,
		"WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	
	//Trace.Println("Test Trace log")
    Info.Println("Test Info log")
    Warn.Println("Test Warn log")
    Error.Println("Test Error log")
}

