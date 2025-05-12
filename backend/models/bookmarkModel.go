package models

type Bookmark struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Pinned     bool   `json:"pinned"`
	FaviconURL string `json:"favicon_url"`
	Category   string `json:"category"`
}
