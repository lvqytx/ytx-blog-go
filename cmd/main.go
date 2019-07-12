package main

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"ytx-blog-go/router"
)

func main() {
	router := &router.Router{}
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
