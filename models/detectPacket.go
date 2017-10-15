package models

// DetectPacket -
type DetectPacket struct {
	Packet     Packet       `bson:",inline"`
	DetectInfo []DetectInfo `bson:"detect_info"`
}
