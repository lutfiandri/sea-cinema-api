package model

import "time"

type Seat struct {
	Name   string `bson:"name"` // 1-64
	UserId string `bson:"user_id"`
}

type Showtime struct {
	Id        string    `bson:"_id"`
	StartTime time.Time `bson:"start_time"`
	EndTime   time.Time `bson:"end_time"`
	Seats     []Seat    `bson:"seats"`
	Price     float64   `bson:"price"`
}

type Movie struct {
	Id          string     `bson:"_id"`
	Title       string     `bson:"title"`
	Description string     `bson:"description"`
	ReleaseDate time.Time  `bson:"release_date"`
	PosterUrl   string     `bson:"poster_url"`
	AgeRating   int        `bson:"age_rating"`
	Shows       []Showtime `bson:"shows"`
	CreatedAt   time.Time  `bson:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at"`
}
