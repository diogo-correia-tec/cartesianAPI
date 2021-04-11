package app

import (
	"cartesianAPI/domain"
	"cartesianAPI/dto"
	"cartesianAPI/util"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// GetPoints godoc
// @Summary Get points that are within the distance range
// @Description Get points that are within the distance range of a given coordinate
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.Response
// @Router /points [get]
// @Param request query dto.Request true "Required Parameters to GetPoints"
func GetPoints(w http.ResponseWriter, r *http.Request) {

	request, err := validateRequest(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	responseList, err := domain.GetPointsWithinDistance(util.CoordinateList, request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, responseList)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func validateRequest(r *http.Request) (dto.Request, error) {

	paramX := r.URL.Query().Get("x")
	paramY := r.URL.Query().Get("y")
	paramDistance := r.URL.Query().Get("distance")

	if len(paramX) <= 0 || len(paramY) <= 0 || len(paramDistance) <= 0 {
		return dto.Request{}, errors.New("no parameters informed for x, y or distance")
	}

	x, errX := strconv.ParseInt(paramX, 10, 64)
	if errX != nil {
		return dto.Request{}, errX
	}

	y, errY := strconv.ParseInt(paramY, 10, 64)
	if errY != nil {
		return dto.Request{}, errY
	}

	distance, errD := strconv.ParseInt(paramDistance, 10, 64)
	if errD != nil {
		return dto.Request{}, errD
	}

	var request dto.Request

	request.X = x
	request.Y = y
	request.Distance = distance

	return request, nil
}
