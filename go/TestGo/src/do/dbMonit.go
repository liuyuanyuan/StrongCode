package do
import (
	"fmt"
	"time"
)

func getDatSize()([]DatSize){
	var sql = "SELECT datname, pg_database_size(oid)/(1024*1024) FROM pg_database WHERE datname NOT IN('template1', 'template0')"
    rows, err := query(sql)
    checkErr(err)

    var datas []DatSize   
    for rows.Next() {
        var datname string
        var datsize int64
        err = rows.Scan(&datname, &datsize)
        //fmt.Println("datname = ", datname, "datsize = ", datsize)
        checkErr(err)
        d := DatSize{datname, datsize, time.Now()}
        datas = append(datas, d)
    }
    
    fmt.Println("Return: size=", len(datas))
    return datas
}

