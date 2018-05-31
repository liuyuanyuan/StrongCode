package main

import (
	"fmt"
	"io"
	"lib"
	"log"
	"net/http"
	
     "database/sql"
     _ "github.com/lib/pq"
)

const num = 1989


/*
type Person struct {
	age  int
	name string
}
*/

var db *sql.DB

func sqlOpen() {
    var err error
    db, err = sql.Open("postgres", "port=5433 user=postgres password=postgres dbname=postgres sslmode=disable")
    // sslmode就是安全验证模式
    // 还可以是这种方式打开
    // db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
    checkErr(err)
}

func sqlInsert() {
    //插入数据
    stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
    checkErr(err)
 
    res, err := stmt.Exec("ficow", "软件开发部门", "2017-03-09")
    //这里的三个参数就是对应上面的$1,$2,$3了 
    checkErr(err)
 
    affect, err := res.RowsAffected()
    checkErr(err)
 
    fmt.Println("rows affect:", affect)
}
func sqlDelete() {
    //删除数据
    stmt, err := db.Prepare("delete from userinfo where uid=$1")
    checkErr(err)
 
    res, err := stmt.Exec(1)
    checkErr(err)
 
    affect, err := res.RowsAffected()
    checkErr(err)
 
    fmt.Println("rows affect:", affect)
}
func sqlSelect() {
    //查询数据
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)
 
    println("-----------")
    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println("uid = ", uid, "\nname = ", username, "\ndep = ", department, "\ncreated = ", created, "\n-----------")
    }
}
func sqlUpdate() {
    //更新数据
    stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
    checkErr(err)
 
    res, err := stmt.Exec("ficow", 1)
    checkErr(err)
 
    affect, err := res.RowsAffected()
    checkErr(err)
 
    fmt.Println("rows affect:", affect)
}
func sqlClose() {
    db.Close()
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main11() {
	 sep := "----------\n"
    sqlOpen()
    println(sep, "*sqlOpen")
    
    rows, err := db.Query("SELECT 123")
    checkErr(err)
    println("-----------")
    for rows.Next() {
    	var uid int
    	err = rows.Scan(&uid)
    	checkErr(err)
    	fmt.Println("***********", uid)
    }
  
 
    sqlSelect()
    println(sep, "*sqlSelect")
 
    sqlInsert()
    sqlSelect()
    println(sep, "*sqlInsert")
 
    sqlUpdate()
    sqlSelect()
    println(sep, "*sqlUpdate")
 
    sqlDelete()
    sqlSelect()
    println(sep, "*sqlDelete")
 
    sqlClose()
    println(sep, "*sqlClose")

	
	
	//same to System.out.println
	fmt.Println(num, "Hello World!", lib.GetTime())

	//struct
	p := lib.Person{
		28,
		"Yuan",
	}
	fmt.Println("Name:", p.Name, ", Age:", p.Age)

	//map
	ages := make(map[string]int)
	ages["linday"] = 20
	ages["michael"] = 30

	fmt.Println(ages["michael"])

	for name, age := range ages {
		fmt.Println("name:", name, ",age:", age)
	}

	delete(ages, "michael")

	for name, age := range ages {
		fmt.Println("name:", name, ",age:", age)
	}
}
