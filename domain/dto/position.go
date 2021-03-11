package dto

//swagger:response
type Position struct {
	// X position of the spaceship
	X float64 `json:"x"`
	// Y position of the spaceship
	Y float64 `json:"y"`
}
