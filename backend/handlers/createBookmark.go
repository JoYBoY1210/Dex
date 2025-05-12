package handlers

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"log"
	"net/http"

	db "github.com/JoYBoY12/Dex/database"
	"github.com/JoYBoY12/Dex/models"
)

func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var bk = models.Bookmark{}
	err := json.NewDecoder(r.Body).Decode(&bk)
	if err != nil || bk.URL == "" {
		http.Error(w, "Invalid bookmark", http.StatusBadRequest)
		return
	}
	favicon := fmt.Sprintf("https://www.google.com/s2/favicons?sz=64&domain_url=%s", bk.URL)
	sql, err := db.Db.Prepare("INSERT INTO bookmarks (title,url,pinned,favicon_url,category) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := sql.Exec(bk.Title, bk.URL, false, favicon, bk.Category)
	if err != nil {
		log.Fatal(err)
		return
	}

	id, _ := result.LastInsertId()
	bk.ID = int(id)
	bk.Pinned = false
	bk.FaviconURL = favicon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bk)
}
