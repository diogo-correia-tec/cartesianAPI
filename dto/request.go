package dto

// swagger:parameters listBars
type Request struct {
	X        int64 `json:"x"`
	Y        int64 `json:"y"`
	Distance int64 `json:"distance"`
}
