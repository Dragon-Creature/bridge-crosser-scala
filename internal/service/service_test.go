package service

import (
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateCrossing(t *testing.T) {
	response, err := CalculateCrossing(model.CrossingRequest{
		Bridges: []model.Bridge{
			{
				LengthInFeet: 100,
				Hikers: []model.Hiker{
					{
						ID:                 uuid.New().String(),
						SpeedFeetInMinutes: 10,
					},
					{
						ID:                 uuid.New().String(),
						SpeedFeetInMinutes: 100,
					},
					{
						ID:                 uuid.New().String(),
						SpeedFeetInMinutes: 50,
					},
					{
						ID:                 uuid.New().String(),
						SpeedFeetInMinutes: 20,
					},
				},
			},
			{
				LengthInFeet: 250,
				Hikers: []model.Hiker{
					{
						ID:                 uuid.New().String(),
						SpeedFeetInMinutes: 2.5,
					},
				},
			},
			{
				LengthInFeet: 150,
				Hikers: []model.Hiker{
					{
						ID:                 uuid.New().String(),
						SpeedFeetInMinutes: 25,
					},
					{
						ID:                 uuid.New().String(),
						SpeedFeetInMinutes: 15,
					},
				},
			},
		},
	})
	require.NoError(t, err)

	const float64EqualityThreshold = 1e-9

	expected := model.CrossingResponse{
		TotalTravelTime: 278,
		BridgeTimeTravel: []float64{
			19,
			150,
			109,
		},
	}
	assert.InDelta(t, expected.TotalTravelTime, response.TotalTravelTime, float64EqualityThreshold)
	assert.Equal(t, len(expected.BridgeTimeTravel), len(response.BridgeTimeTravel))
	for i, _ := range response.BridgeTimeTravel {
		assert.InDelta(t, expected.BridgeTimeTravel[i], response.BridgeTimeTravel[i], float64EqualityThreshold)
	}
}
