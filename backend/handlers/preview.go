package handlers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetBookmarkPreview(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	pageData, err := fetchPageMetadata(url)
	if err != nil {
		http.Error(w, "Failed to fetch metadata", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pageData)

}

func fetchPageMetadata(url string) (map[string]string, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to fetch the page")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	metaData := map[string]string{
		"title":       doc.Find("head title").Text(),
		"favicon":     doc.Find("link[rel='icon']").AttrOr("href", ""),
		"description": doc.Find("meta[name='description']").AttrOr("content", ""),
	}

	return metaData, nil

}
