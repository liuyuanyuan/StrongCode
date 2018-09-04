package do

import(
	"time"
	"database/sql"
)

//db
type DatSize struct {
	Current     string    `json:"current"` //time.Time
	
    DatName     string    `json:"datname"`
    DatSize     int64     `json:"datsize"`
    //Completed bool   `json:"completed"`
}
type DatSession struct{
	Current      string     `json:"current"` //time.Time
	
	Pid          int        `json:"pid"`
    DatName      sql.NullString    `json:"datname"`
    UserName     sql.NullString    `json:"username"`
    AppName      sql.NullString    `json:"appname"`
    ClientAddr   sql.NullString    `json:"clientaddr"`
    ClientPort   sql.NullInt64     `json:"clientport"`
    State        sql.NullString    `json:"status"`
    Query        sql.NullString    `json:"query"` 
    BackendStart time.Time  `json:"backendStart"`
}


//os
type OSMem struct {
	Current string `json:"current"` //time.Time
	// /proc/meminfo
    Total   uint64  `json:"all"`
    Free    uint64  `json:"free"`
    Used    uint64  `json:"used"`
    //Self  uint64  `json:"self"`
}
type OSCpu struct {
	Current string `json:"current"` //time.Time
	// /proc/stat
	CPU       string
	User      float64
	Nice      float64
	System    float64
	Idle      float64
	Iowait    float64
	Irq       float64
	Softirq   float64
	Stolen    float64
	
	Steal     float64
	Guest     float64
	GuestNice float64
}
type OSLoadAvg struct {
	Current string `json:"current"` //time.Time
    // /proc/loadavg
	Load1   float64 `json:"Load1"`  /*load avg of last 1 min*/
	Load5   float64 `json:"Load5"`  /*load avg of last 5 min*/
	Load15  float64 `json:"Load15"` /*load avg of last 15 min*/
}



//test
type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}
type Todos []Todo


