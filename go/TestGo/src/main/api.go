package main

//简单的 JSON Restful API 演示(服务端) 
//ref: https://studygolang.com/articles/3603
import (
	"log"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}


func main0() {
	http.HandleFunc("/", JsonHandler)
	err := http.ListenAndServe("192.168.100.172:8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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

