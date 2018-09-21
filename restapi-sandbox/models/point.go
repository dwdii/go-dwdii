package models

type Point struct {
	UserId    string  `json:"UserId"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
}
