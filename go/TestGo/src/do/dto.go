package do

import(
	"time"
	"database/sql"
)

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


type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}
type Todos []Todo


