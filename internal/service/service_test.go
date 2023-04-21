package service

import (
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"
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
						SpeedFeetInMinutes: 10,
					},
					{
						SpeedFeetInMinutes: 100,
					},
					{
						SpeedFeetInMinutes: 50,
					},
					{
						SpeedFeetInMinutes: 20,
					},
				},
			},
			{
				LengthInFeet: 250,
				Hikers: []model.Hiker{
					{
						SpeedFeetInMinutes: 2.5,
					},
				},
			},
			{
				LengthInFeet: 150,
				Hikers: []model.Hiker{
					{
						SpeedFeetInMinutes: 25,
					},
					{
						SpeedFeetInMinutes: 15,
					},
				},
			},
		},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
}
