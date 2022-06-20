package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func callback(w http.ResponseWriter, req *http.Request) {
	var p map[string]string
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(p["link"])
}

func main() {

	http.HandleFunc("/callback", callback)

	http.ListenAndServe(":8090", nil)
}
