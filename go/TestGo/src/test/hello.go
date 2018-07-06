package main


import (
	"fmt"
	"log"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}



func main1(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "URL.Path = %q\n", r.URL.Path)
	
}




