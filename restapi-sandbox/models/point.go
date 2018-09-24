package models

type Point struct {
	UserId    string  `json:"userid"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
	Timestamp int64   `json:"timestamp"`
}
