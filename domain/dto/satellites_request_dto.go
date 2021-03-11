package dto

//Topsecret request model,
type SatelliteRequest struct {
	//Satellites List
	Satellites []Satellite `json:"satellites,ommitempty"`
}
