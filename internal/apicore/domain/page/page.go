package page

import "time"

// Page .
type Page struct {
	Id          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Author      string    `json:"author,omitempty"`
	Thumbnail   string    `json:"thumbnail,omitempty"`
	Views       int32     `json:"views,omitempty"`
	Content     string    `json:"content,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func NewPage() (*Page, error) {
	return &Page{}, nil
}
