package service

import (
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateCrossing(t *testing.T) {
	response := CalculateCrossing(model.CrossingRequest{
		Bridges: []model.Bridge{
			{
				LengthInFeet: 100,
				Hikers: []model.Hiker{
					{
						ID:                 "9dfc18ae-e8fd-4d9a-91be-11964319fe87",
						SpeedFeetInMinutes: 10,
					},
					{
						ID:                 "68f4e36f-efdd-4c0e-b8c8-6000b3ce6c4c",
						SpeedFeetInMinutes: 100,
					},
					{
						ID:                 "e331f9be-c9c1-4503-9ccb-8f787707edde",
						SpeedFeetInMinutes: 50,
					},
					{
						ID:                 "86b7e86b-650a-4ce1-aff1-fa2742521b4a",
						SpeedFeetInMinutes: 20,
					},
				},
			},
			{
				LengthInFeet: 250,
				Hikers: []model.Hiker{
					{
						ID:                 "0ed010b3-505c-4933-9e77-39332c28cd81",
						SpeedFeetInMinutes: 2.5,
					},
				},
			},
			{
				LengthInFeet: 150,
				Hikers: []model.Hiker{
					{
						ID:                 "5ed6da13-f154-4872-8433-fb3c9e843fe5",
						SpeedFeetInMinutes: 25,
					},
					{
						ID:                 "1c2a6ab9-e610-42a8-9130-c4d7e90c097e",
						SpeedFeetInMinutes: 15,
					},
				},
			},
		},
	})
	const float64EqualityThreshold = 1e-9

	expected := model.CrossingResponse{
		TotalTravelTime: 278,
		BridgeResults: []model.BridgeResult{
			{
				LengthInFeet:          100,
				NumberOfHikersCrossed: 4,
				IDOfHikers: []string{"9dfc18ae-e8fd-4d9a-91be-11964319fe87",
					"68f4e36f-efdd-4c0e-b8c8-6000b3ce6c4c", "e331f9be-c9c1-4503-9ccb-8f787707edde",
					"86b7e86b-650a-4ce1-aff1-fa2742521b4a",
				},
				TotalTravelTime: 19,
			},
			{
				LengthInFeet:          250,
				NumberOfHikersCrossed: 5,
				IDOfHikers: []string{"9dfc18ae-e8fd-4d9a-91be-11964319fe87",
					"68f4e36f-efdd-4c0e-b8c8-6000b3ce6c4c", "e331f9be-c9c1-4503-9ccb-8f787707edde",
					"86b7e86b-650a-4ce1-aff1-fa2742521b4a", "0ed010b3-505c-4933-9e77-39332c28cd81",
				},
				TotalTravelTime: 150,
			},
			{
				LengthInFeet:          150,
				NumberOfHikersCrossed: 7,
				IDOfHikers: []string{"9dfc18ae-e8fd-4d9a-91be-11964319fe87",
					"68f4e36f-efdd-4c0e-b8c8-6000b3ce6c4c", "e331f9be-c9c1-4503-9ccb-8f787707edde",
					"86b7e86b-650a-4ce1-aff1-fa2742521b4a", "0ed010b3-505c-4933-9e77-39332c28cd81",
					"5ed6da13-f154-4872-8433-fb3c9e843fe5", "1c2a6ab9-e610-42a8-9130-c4d7e90c097e",
				},
				TotalTravelTime: 109,
			},
		},
	}
	assert.InDelta(t, expected.TotalTravelTime, response.TotalTravelTime, float64EqualityThreshold)
	assert.Equal(t, len(expected.BridgeResults), len(response.BridgeResults))
	for i, _ := range response.BridgeResults {
		assert.InDelta(t, expected.BridgeResults[i].TotalTravelTime, response.BridgeResults[i].TotalTravelTime, float64EqualityThreshold)
		assert.InDelta(t, expected.BridgeResults[i].LengthInFeet, response.BridgeResults[i].LengthInFeet, float64EqualityThreshold)
		assert.Equal(t, expected.BridgeResults[i].NumberOfHikersCrossed, response.BridgeResults[i].NumberOfHikersCrossed)
		assert.ElementsMatch(t, expected.BridgeResults[i].IDOfHikers, response.BridgeResults[i].IDOfHikers)
	}
}
