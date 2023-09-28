package model

type TimeTrack struct {
	CreatedAt int `json:"created_at" bson:"created_at"`
	UpdatedAt int `json:"updated_at" bson:"updated_at"`
	DeletedAt int `json:"deleted_at" bson:"deleted_at"`
}
