package model

type CrossingRequest struct {
	Bridges []Bridge `json:"bridges"`
}

type CrossingResponse struct {
	TotalTravelTime  float64   `json:"total_travel_time"`
	BridgeTimeTravel []float64 `json:"bridge_time_travel"`
}

type Bridge struct {
	LengthInFeet float64 `json:"length_in_feet"`
	Hikers       []Hiker `json:"hikers"`
}

type Hiker struct {
	ID                 string  `json:"id"`
	SpeedFeetInMinutes float64 `json:"speed_feet_in_minutes"`
}

type HikerResponse struct {
	ID                  string    `json:"id"`
	TravelTimeTotal     float64   `json:"travel_time_total"`
	TravelTimePerBridge []float64 `json:"travel_time_per_bridge"`
}
