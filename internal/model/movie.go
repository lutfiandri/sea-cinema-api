package model

import "time"

type MovieShowSeat struct {
	Name   string `bson:"name"` // 1-64
	UserId string `bson:"user_id"`
}

type MovieShow struct {
	Id        string          `bson:"_id"`
	StartedAt time.Time       `bson:"started_at"`
	EndedAt   time.Time       `bson:"ended_at"`
	Seats     []MovieShowSeat `bson:"seats"`
	Price     float64         `bson:"price"`
}

type Movie struct {
	Id          string      `bson:"_id"`
	Title       string      `bson:"title"`
	Description string      `bson:"description"`
	ReleaseDate time.Time   `bson:"release_date"`
	PosterUrl   string      `bson:"poster_url"`
	AgeRating   int         `bson:"age_rating"`
	Shows       []MovieShow `bson:"shows"`
	CreatedAt   time.Time   `bson:"created_at"`
	UpdatedAt   time.Time   `bson:"updated_at"`
}
