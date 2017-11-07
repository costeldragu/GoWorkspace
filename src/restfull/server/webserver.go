package webserver

import (
	"net/http"
	"fmt"
)

const (
	Port = ":8080"
)

func ServeDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now "
	fmt.Fprintln(w,response)
}