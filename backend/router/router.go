package router

import (
	"github.com/JoYBoY12/Dex/handlers"
	"github.com/gorilla/mux"
)

func SetUpNewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create", handlers.CreateBookmark).Methods("POST")
	r.HandleFunc("/bookmarks", handlers.GetBookmark).Methods("GET")
	r.HandleFunc("/pin/bookmarks/{id}", handlers.TogglePinBookmark).Methods("PATCH")
	r.HandleFunc("/delete/bookmarks/{id}", handlers.DeleteBookmark).Methods("DELETE")
	r.HandleFunc("/pinned/bookmarks", handlers.GetPinnedBookmark).Methods("GET")
	r.HandleFunc("/bookmarks/preview", handlers.GetBookmarkPreview).Methods("GET")

	return r
}
