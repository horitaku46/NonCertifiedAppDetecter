package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Packet -
type Packet struct {
	ID         bson.ObjectId `bson:"_id"`
	Timestamp  time.Time     `bson:"timestamp"`
	URL        string        `bson:"url"`
	HTTPMethod string        `bson:"http_method"`
	Host       string        `bson:"host"`
	Packet     string        `bson:"packet"`
}
