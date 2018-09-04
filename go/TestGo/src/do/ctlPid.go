package do

import (
	"os"
	"io/ioutil"
	"strconv"
	"syscall"
)

func SavePid(){
	Info.Println("Current pid=", strconv.Itoa(os.Getpid()))
	
	fileName := "goagent.pid"
	
	f, err := os.OpenFile(fileName, os.O_RDWR, 0600)    
    Error.Println(err)
    if f == nil {
    	f,err := os.Create(fileName)
	    defer f.Close()
	    Error.Println(err)
	    
        _,err = f.WriteString(strconv.Itoa(os.Getpid()))
        Info.Println("GoAgent is starting ...")
        checkErr(err)
    } else {
	     defer f.Close()
	     contentByte,err := ioutil.ReadAll(f)
         Error.Println(err)		 
		 
		 if contentByte == nil{
		 	_,err = f.WriteString(strconv.Itoa(os.Getpid()))
		 	Info.Println("GoAgent is starting ...")
	        checkErr(err)
		 } else {
			Info.Println("Old pid=",string(contentByte))
			oldPid, err := strconv.Atoi(string(contentByte))
			Error.Println(err)
			op, err :=  os.FindProcess(oldPid)//always returns for unix
			if err != nil {
				Error.Println(err)
				 checkErr(err)
			} else {
				 err = op.Signal(syscall.Signal(0))
				 if err != nil{
				 	f.Truncate(0) //empty
				 	f.Seek(0,0) //set offset for the next read or write
				 	_,err = f.WriteString(strconv.Itoa(os.Getpid()))
				 	Info.Println("GoAgent is not running, so start now...")
				 } else {
				 	Info.Println("GoAgent is running, do nothing and return.")
				 }
			}
			
			/*
			//LINUX
			oldPidFile, err := os.Open("/proc/" + string(contentByte) + "/stat")
			Error.Println(err)
			Info.Println("oldPidFile=", oldPidFile)
			if oldPidFile == nil {
				f.Truncate(0)
				f.Seek(0,0)
				checkErr(err)
				_,err = f.WriteString(strconv.Itoa(os.Getpid()))
		        checkErr(err)
			}
			*/
			/*
			//Win
			oldPid, err := strconv.Atoi(string(contentByte))
			Error.Println(err)
			op, err :=  os.FindProcess(oldPid)//always returns for unix
			if err == nil {
			    f.Truncate(0)
				f.Seek(0,0)
				_,err = f.WriteString(strconv.Itoa(os.Getpid()))
		        checkErr(err)
			}
			*/
		 }
    }
}