package do
import (
	//"time"
	//"log"
	"fmt"
	
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
    todos := Todos{
	    Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func GetDatSize(w http.ResponseWriter, r *http.Request){
    datas := datSizeMap[len(datSizeMap)-1] //getDatSize()
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    err := json.NewEncoder(w).Encode(datas)
    checkErr(err)
    
    /*if err := json.NewEncoder(w).Encode(datas); err != nil {
        panic(err)
    }*/
}

func GetDatSession(w http.ResponseWriter, r *http.Request){
    datas := datSessionMap[len(datSessionMap)-1] //getDatSession()
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    err := json.NewEncoder(w).Encode(datas)
    checkErr(err)
    
    /*if err := json.NewEncoder(w).Encode(datas); err != nil {
        panic(err)
    }*/
}

