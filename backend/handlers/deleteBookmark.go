package handlers

import (

	// "fmt"
	// "encoding/json"
	// "log"
	"net/http"

	db "github.com/JoYBoY12/Dex/database"
	// "github.com/JoYBoY12/Dex/models"
	"github.com/gorilla/mux"
)

func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	sql, err := db.Db.Prepare("DELETE FROM bookmarks WHERE id = ?")
	if err != nil {
		http.Error(w, "Bookmark not found", http.StatusNotFound)
		return
	}
	_, err = sql.Exec(id)
	if err != nil {
		http.Error(w, "Failed to delete bookmark", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}
