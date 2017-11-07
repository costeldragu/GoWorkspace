package main

import (
	"log"
	"com.enva/src/restufull/server"
	"net/http"
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
	http.HandleFunc("/",webserver.ServeDynamic )
	http.ListenAndServe(webserver.Port,nil)
}
