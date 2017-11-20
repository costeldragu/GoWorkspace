package main

import (
	"log"
	"net/http"
	"restfull/server"
	"github.com/gorilla/mux"
)

const (
	Port = ":8080"
)

/**
 * Initialization of the program
 */
func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("Server has start")
	routes := mux.NewRouter()
	routes.HandleFunc("/api/{guid:[0-9a-zA\\-]+}",webserver.ServeDynamic ).Methods("GET")
	routes.HandleFunc("/pages/{id:[0-9]+}", webserver.PageHandler).Methods("GET")
	routes.HandleFunc("/{pageId:[0-9a-zA\\-]+}", webserver.ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(Port,nil)
}
