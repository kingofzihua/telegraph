package biz

import "time"

type Page struct {
	// Path to the page.
	Path string `json:"path,omitempty"`
	// Title of the page.
	Title string `json:"title,omitempty"`
	// Description of the page.
	Description string `json:"description,omitempty"`
	// Optional. Name of the author, displayed below the title.
	Author string `json:"author,omitempty"`
	// Optional. Image URL of the page.
	Thumbnail string `json:"thumbnail,omitempty"`
	// Number of page views for the page.
	Views int32 `json:"views,omitempty"`
	// Optional. Content of the page.
	Content string `json:"content,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
