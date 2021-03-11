package dto

type Satellite struct {
	Name     string   `json:"name,omitempty"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}
