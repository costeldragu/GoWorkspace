package webserver

import (
	"net/http"
	"time"
	"fmt"
)

const (
	Port = ":8080"
)

func ServeDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()
	fmt.Fprintln(w,response)
}