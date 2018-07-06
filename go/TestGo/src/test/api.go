package main

//JSON Restful API example(server side) 
//ref: https://studygolang.com/articles/3603
//run Restful api by: go run api.go
import (
	"log"
	"fmt"
	"time"
	
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux" //must under Go/src or GoWorkspace/src
)

//model
type Item struct {
	Seq    int
	Result map[string]int
}
type Message struct {
	Dept    string
	Subject string
	Time    int64
	Detail  []Item
}

//api
func getJsonMsg() ([]byte, error) {
	pass := make(map[string]int)
	pass["x"] = 50
	pass["c"] = 60
	item1 := Item{100, pass}

	reject := make(map[string]int)
	reject["l"] = 11
	reject["d"] = 20
	item2 := Item{200, reject}

	detail := []Item{item1, item2}
	
	m := Message{"IT", "KPI", time.Now().Unix(), detail}
	return json.MarshalIndent(m, "", "")
}


func JsonHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := getJsonMsg()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(resp))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}


type Todo struct {
    Name        string
    Completed   int64
    Due         time.Time
}
type Todos []Todo
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	//return normal 
    //fmt.Fprintln(w, "Todo Index!") 
    
    //return json
    todos := Todos{
    Todo{Name: "Write presentation"},
    Todo{Name: "Host meetup"},
    }
    json.NewEncoder(w).Encode(todos) 
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}




func main() {
	//增加路径分发功能	
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)                    //http://192.168.100.172:8085
    router.HandleFunc("/todos", TodoIndex)           //http://192.168.100.172:8085/todos
    router.HandleFunc("/todos/{todoId}", TodoShow)   //http://192.168.100.172:8085/todos/3	
	router.HandleFunc("/jsonmsg", JsonHandler)       //http://192.168.100.172:8085/jsonmsg
	
    log.Fatal(http.ListenAndServe(":8085", router))  //http://host:8085
}

func main0() {
	//simple start server
	http.HandleFunc("/", JsonHandler)
	err := http.ListenAndServe("192.168.100.172:8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
