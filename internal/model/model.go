package model

type CrossingRequest struct {
	Bridges []Bridge `json:"bridges"`
	Hikers  []Hiker  `json:"hikers"`
}

type CrossingResponse struct {
}

type Bridge struct {
	LengthInFeet int `json:"length_in_feet"`
}

type Hiker struct {
	SpeedFeetInMinutes float64 `json:"speed_feet_in_minutes"`
}
