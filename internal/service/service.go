package service

import (
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"
	"sort"
)

func CalculateCrossing(request model.CrossingRequest) (*model.CrossingResponse, error) {
	//It doesn't matter if they go all the way back or one bridge at a time.
	//Always send the fastest as company when someone crosses the bridge, to increase round trip back.
	totalMinutesOfTravel := float64(0)
	for i := 0; i < len(request.Bridges); i++ {
		hikers := request.Bridges[i].Hikers
		bridgeLengthInFeet := request.Bridges[i].LengthInFeet
		sort.Slice(hikers, func(j, k int) bool {
			if hikers[j].SpeedFeetInMinutes > hikers[k].SpeedFeetInMinutes {
				return true
			}
			return false
		})
		for len(hikers) > 0 {
			fastest := hikers[0]
			companion := hikers[1]
			if len(request.Bridges) != i+1 {
				totalMinutesOfTravel += bridgeLengthInFeet / companion.SpeedFeetInMinutes
				request.Bridges[i+1].Hikers = append(request.Bridges[i+1].Hikers, companion)
				hikers = remove(hikers, 1)
				if len(hikers) <= 1 {
					request.Bridges[i+1].Hikers = append(request.Bridges[i+1].Hikers, fastest)
					hikers = remove(hikers, 0)
				}
			} else if len(request.Bridges) == i+1 {
				totalMinutesOfTravel += bridgeLengthInFeet / companion.SpeedFeetInMinutes
				hikers = remove(hikers, 1)
				if len(hikers) <= 1 {
					hikers = remove(hikers, 0)
				}
			}
		}
		request.Bridges[i].Hikers = hikers
	}
	return nil, nil
}

func remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
