package main

import (
	"log"
	"net/http"

	"github.com/teimurjan/go-plain-server/handlers/page_handler"
)

func main() {
	http.HandleFunc("/view/", page_handler.ViewHandler)
	http.HandleFunc("/", page_handler.ViewAllHandler)
	http.HandleFunc("/edit/", page_handler.EditHandler)
	http.HandleFunc("/save/", page_handler.SaveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
