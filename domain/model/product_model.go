package model

type ProductModel struct {
	ID          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Price       float32   `bson:"price"`
	Stock       int64     `bson:"stock"`
	Images      []string  `bson:"images"`
	TimeTrack   TimeTrack `bson:"time_track"`
}
