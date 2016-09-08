package db

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Item struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Name  string        `json:"name" bson:"name"`
		Image string        `json:"image" bson:"image"`
		Price string        `json:"price" bson:"price"`
		Size  string        `json:"size" bson:"size"`
	}
)
