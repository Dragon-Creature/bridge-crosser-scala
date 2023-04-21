package model

type CrossingRequest struct {
	Bridges []Bridge `json:"bridges"`
}

type CrossingResponse struct {
	TotalTravelTime float64        `json:"total_travel_time"`
	BridgeResults   []BridgeResult `json:"bridge_results"`
}

type BridgeResult struct {
	LengthInFeet          float64  `json:"length_in_feet"`
	NumberOfHikersCrossed int      `json:"number_of_hikers_crossed"`
	IDOfHikers            []string `json:"id_of_hikers"`
	TotalTravelTime       float64  `json:"total_travel_time"`
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
