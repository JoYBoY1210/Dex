package handlers

import (

	// "fmt"

	"log"
	"net/http"

	db "github.com/JoYBoY12/Dex/database"
	"github.com/gorilla/mux"
)

func TogglePinBookmark(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	sql, err := db.Db.Prepare("UPDATE bookmarks SET pinned = NOT pinned WHERE id = ?")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = sql.Exec(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
