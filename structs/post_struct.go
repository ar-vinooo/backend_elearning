package structs

import "time"

type Post struct {
	GuruID       int       `json:"guru_id"`
	GuruName     string    `json:"guru_name"`
	GuruNip      string    `json:"guru_nip"`
	KelasID      int       `json:"kelas_id"`
	KelasName    string    `json:"kelas_name"`
	PostID       int       `json:"post_id"`
	Type         string    `json:"type"`
	DiscussionID int       `json:"discussion_id"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
}
