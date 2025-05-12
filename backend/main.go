package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/JoYBoY12/Dex/database"
	"github.com/JoYBoY12/Dex/router"

	"github.com/rs/cors"
)

func main() {
	r := router.SetUpNewRouter()
	db.InitDB()
	defer db.CloseDB()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})
	handler := c.Handler(r)
	fmt.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal("failed to start ", err)

	}

}
