package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() 
	const port string =":8000"
	router.HandleFunc("/",func(response http.ResponseWriter, request *http.Request){
          fmt.Fprintln(response,"Up and runnuing....")
	})
	log.Println("Server listing on port",port)
	log.Fatalln(http.ListenAndServe(port,router))

}