package structs

type Guru struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Nip    string `json:"nip"`
	Gender string `json:"gender"`
}
