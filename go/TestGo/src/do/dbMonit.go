package do
import (
	"fmt"
	"time"
)

func getDatSize()([]DatSize){
	var sql = `SELECT datname, pg_database_size(oid)/(1024*1024) 
	FROM pg_database 
	WHERE datname NOT IN('template1', 'template0')`
    rows, err := query(sql)
    checkErr(err)

    var datas []DatSize   
    for rows.Next() {
        var datname string
        var datsize int64
        err = rows.Scan(&datname, &datsize)
        //fmt.Println("datname = ", datname, "datsize = ", datsize)
        checkErr(err)
        d := DatSize{time.Now(), datname, datsize}
        datas = append(datas, d)
    }
    
    fmt.Println("Return: size=", len(datas))
    return datas
}


func getDatSession()([]DatSession){
	var sql = `SELECT  pid, datname, usename, application_name, 
	client_addr, client_port, state, query, backend_start
	FROM pg_stat_activity`
	rows, err := query(sql)
    checkErr(err)
    
    var datas []DatSession
    for rows.Next() {
    	var data DatSession
    	data.Current = time.Now()
        err = rows.Scan( &data.Pid, &data.DatName, &data.UserName, &data.AppName,
	        &data.ClientAddr, &data.ClientPort, &data.State, &data.Query, &data.BackendStart)
	    checkErr(err)
	    datas = append(datas, data)        
    }
    
    fmt.Println("Return: size=",len(datas))
	return datas
}
	




