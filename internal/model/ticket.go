package model

import "time"

type TicketLog struct {
	Id          string    `bson:"_id"`
	UserId      string    `bson:"user_id"`
	MovieShowId string    `bson:"movie_id"`
	Action      string    `bson:"action"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
