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

	OnSale struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Name   string        `json:"name" bson:"name"`
		Image  string        `json:"image" bson:"image"`
		Price  string        `json:"price" bson:"price"`
		SPrice string        `json:"sprice" bson:"sprice"`
		Size   string        `json:"size" bson:"size"`
	}

	Address struct {
		Id      bson.ObjectId `json:"id" bson:"_id"`
		Name    string        `json:"name" bson:"name"`
		Mobile  string        `json:"mobile" bson:"mobile"`
		Region  string        `json:"region" bson:"region"`
		Address string        `json:"address" bson:"address"`
	}
)
