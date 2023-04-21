package model

type CrossingRequest struct {
	Bridges []Bridge `json:"bridges"`
}

type CrossingResponse struct {
}

type Bridge struct {
	LengthInFeet float64 `json:"length_in_feet"`
	Hikers       []Hiker `json:"hikers"`
}

type Hiker struct {
	SpeedFeetInMinutes float64 `json:"speed_feet_in_minutes"`
}
