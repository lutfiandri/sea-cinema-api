package contract

import "time"

type GetProfileResponse struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	BirthDate time.Time `json:"birth_date"`
}

type TopUpBalanceRequest struct {
	Balance float64 `json:""`
}

type TopUpBalanceResponse struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}
