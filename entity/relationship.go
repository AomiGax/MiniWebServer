package entity

// Relationship struct
type Relationship struct {
	ID      int    `json:"id"`
	UID     int    `json:"uid"`
	OtherID int    `json:"other_id"`
	State   string `json:"state"`
	Type    string "user"
}
