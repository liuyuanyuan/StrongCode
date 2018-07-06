package do
import(
	"time"
)

type DatSize struct {
    DatName     string    `json:"datname"`
    DatSize     int64     `json:"datsize"`
    CurrentTime time.Time `json:"due"`
    //Completed    bool   `json:"completed"`
}

//type DatSizeArray []DatSize