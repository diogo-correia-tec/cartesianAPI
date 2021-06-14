package domain

import (
	"errors"
	"math"
	"sort"

	"github.com/diogo-correia-tec/cartesianAPI/dto"
)

// Coordinate is the base data been validated
type Coordinate struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

//GetPointsWithinDistance check wich points are within the distance range
func GetPointsWithinDistance(list []Coordinate, r dto.Request) ([]dto.Response, error) {

	var responseList []dto.Response

	for _, point := range list {

		reqCoordinate := Coordinate{r.X, r.Y}

		resp, isWithin := validateManhattanDistance(reqCoordinate, point, r.Distance)

		if isWithin {
			responseList = append(responseList, resp)
		}
	}

	if responseList == nil || len(responseList) <= 0 {
		return nil, errors.New("no points within distance")
	}

	return sortListByDistance(responseList), nil
}

func validateManhattanDistance(p1 Coordinate, p2 Coordinate, distanceToValidate int64) (dto.Response, bool) {
	response := dto.Response{X: p2.X, Y: p2.Y, Distance: 0}

	distance := int64((math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))) //distance(p1,p2) = |x1-x2| + |y1-y2|
	response.Distance = distance

	if distance <= distanceToValidate {
		return response, true
	}
	return response, false
}

func sortListByDistance(responseList []dto.Response) []dto.Response {
	sort.Slice(responseList, func(i, j int) bool {
		return responseList[i].Distance < responseList[j].Distance
	})

	return responseList
}
