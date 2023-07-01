package model

import "time"

type User struct {
	Id        string    `bson:"_id"`
	Username  string    `bson:"username"`
	Password  string    `bson:"password"`
	Name      string    `bson:"name"`
	Balance   float64   `bson:"balance"`
	BirthDate time.Time `bson:"birth_date"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
