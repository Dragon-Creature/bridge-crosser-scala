package model

type CrossingRequest struct {
	Bridges []Bridge `json:"bridges" valid:"required"`
}

type Bridge struct {
	LengthInFeet float64 `json:"length_in_feet" valid:"required,numeric"`
	Hikers       []Hiker `json:"hikers" valid:"required"`
}

type Hiker struct {
	ID                 string  `json:"id" valid:"required,uuid"`
	SpeedFeetInMinutes float64 `json:"speed_feet_in_minutes" valid:"required,numeric"`
}
