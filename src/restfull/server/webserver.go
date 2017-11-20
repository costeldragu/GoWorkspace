package webserver

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)


func ServeDynamic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	response := "ServeDynamic " + pageGUID
	log.Println("Server Dynamic page" + pageGUID)
	fmt.Fprintln(w,response)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	response := "PageHandler " + pageID
	log.Println("Page Handler page" + pageID)
	fmt.Fprintln(w,response)
}

func ServePage(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r);
    pageID := vars["pageId"]
	response := "ServePage " + pageID
	log.Println("Page Handler page" + pageID)
	fmt.Fprintln(w,response)
}