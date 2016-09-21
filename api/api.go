package api

import (
	"github.com/efegurkan/dummyhttp/controller"
	"github.com/efegurkan/dummyhttp/utils"
	"log"
	"net/http"
	"strings"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Api handler started")
	var substr string = r.URL.Path[1:]
	var index int = strings.Index(substr, "/") + 1
	substr = substr[index:]
	file, err := utils.ReadJsonFromFile(substr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(file)
	}
}

func InitializeServer(port string) {
	log.Println("Starting at port: ", port)
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/api/", apiHandler)
	http.ListenAndServe(":"+port, nil)
}
