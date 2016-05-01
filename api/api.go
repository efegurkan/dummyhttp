package api

import (
	"github.com/efegurkan/dummyhttp/utils"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	file, err := utils.ReadJsonFromFile(r.URL.Path[1:])
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
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":"+port, nil)
}
