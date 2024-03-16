package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to homepage")
	fmt.Println("Endpoint hit")
}


func main(){
    http.HandleFunc("/", homepage)
    err :=  http.ListenAndServe(":12000",nil)
    if err != nil{
	    fmt.Println("failure in starting the server:", err)
    }


}
