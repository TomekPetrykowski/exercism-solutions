package darts

import "math"

func Score(x, y float64) int {
	
	// Score table sorted by radius
	scoringTable := []struct {
		radius float64
		points int
	}{
		{radius: 1, points: 10},
		{radius: 5, points: 5},
		{radius: 10, points: 1},
	}

	for _, scoring := range scoringTable {
		distance := math.Sqrt(x*x + y*y)
		
		if distance <= scoring.radius {
			return scoring.points
		}
	}
	
	return 0
}
