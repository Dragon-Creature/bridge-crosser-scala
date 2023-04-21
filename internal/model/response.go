package model

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

type Error struct {
	HttpCode int    `json:"http_code"`
	Message  string `json:"message"`
}
