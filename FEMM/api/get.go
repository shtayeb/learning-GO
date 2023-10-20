package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"shtb.dev/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query()["id"]

	if id != nil { // We will serve only one
		finalId, err := strconv.Atoi(id[0])

		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else { // We will serve all
		json.NewEncoder(w).Encode(data.GetAll())
	}

}
