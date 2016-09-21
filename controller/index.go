package controller

import (
	"github.com/efegurkan/dummyhttp/utils"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Index controller started.")
	file, err := utils.ReadHttpFile("public/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.Write(file)
	}
}
