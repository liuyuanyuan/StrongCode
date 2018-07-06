package do
import(
	"time"
	"database/sql"
)

type DatSize struct {
	Current     time.Time `json:"current"`
	
    DatName     string    `json:"datname"`
    DatSize     int64     `json:"datsize"`
    //Completed bool   `json:"completed"`
}

type DatSession struct{
	Current     time.Time `json:"current"`
	
	Pid          int       `json:"pid"`
    DatName      sql.NullString    `json:"datname"`
    UserName     sql.NullString    `json:"username"`
    AppName      sql.NullString    `json:"appname"`
    ClientAddr   sql.NullString    `json:"clientaddr"`
    ClientPort   sql.NullInt64     `json:"clientport"`
    State        sql.NullString    `json:"status"`
    Query        sql.NullString    `json:"query"` 
    BackendStart time.Time  `json:"backendStart"`
}

