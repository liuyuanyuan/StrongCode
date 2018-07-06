package do

import (
	"fmt"
	//"io"
	//"lib"
	//"log"
	//"net/http"
	
     "database/sql"
     _ "github.com/lib/pq"
)

var db *sql.DB
// init is called prior to main.
func init() {
	sqlOpen()
}
func sqlOpen(){
    var dbUrl = "port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"
    var err error
    db, err = sql.Open("postgres", dbUrl)
    // sslmode就是安全验证模式, 还可以是这种方式打开
    // db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
    checkErr(err)
}

func query(sql string)(*sql.Rows, error) {
	fmt.Println("Enter:", sql)
    rows, err := db.Query(sql)
    return rows, err
}

func sqlClose() {
    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

/*
func TestQuery() {
    rows, err := db.Query("SELECT datname, pg_database_size(oid)/(1024*1024) FROM pg_database")
    checkErr(err)
 
    for rows.Next() {        
        var datname string
        var size int
        err = rows.Scan(&datname, &size)
        checkErr(err)
        fmt.Println("datname = ", datname, "size = ", size)
    }
}

func sqlInsert() {
    //插入数据
    stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
    checkErr(err)
 
	 //这里的三个参数就是对应上面的$1,$2,$3了
    res, err := stmt.Exec("ficow", "软件开发部门", "2017-03-09")
    checkErr(err)
 
    affect, err := res.RowsAffected()
    checkErr(err)
 
    fmt.Println("rows affect:", affect)
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
*/


