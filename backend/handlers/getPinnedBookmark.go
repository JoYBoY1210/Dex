package handlers

import (
	"log"
	"net/http"

	db "github.com/JoYBoY12/Dex/database"
	// "github.com/gorilla/mux"
	"encoding/json"

	"github.com/JoYBoY12/Dex/models"
)

func GetPinnedBookmark(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Db.Query("SELECT id,title,url,pinned,favicon_url,category FROM bookmarks WHERE pinned = 1")
	if err != nil {
		http.Error(w, "Error fetching bookmarks", http.StatusInternalServerError)
		log.Println("Error fetching bookmarks:", err)
		return
	}
	defer rows.Close()
	var bookmarks []models.Bookmark

	for rows.Next() {
		var bk models.Bookmark
		err := rows.Scan(&bk.ID, &bk.Title, &bk.URL, &bk.Pinned, &bk.FaviconURL, &bk.Category)
		if err != nil {
			http.Error(w, "Error scanning bookmarks", http.StatusInternalServerError)
			log.Println("Error scanning bookmarks:", err)
			return
		}
		bookmarks = append(bookmarks, bk)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookmarks)
}
