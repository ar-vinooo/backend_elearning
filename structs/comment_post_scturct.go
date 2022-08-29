package structs

import "time"

type CommentPost struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	UserType   string    `json:"user_type"`
	UserName   string    `json:"user_name"`
	UserNumber string    `json:"user_number"`
	PostID     int       `json:"post_id"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}
