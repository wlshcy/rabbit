package db

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Phone string `json:"phone" bson:"phone"`
	}

	Credential struct {
		Phone string `json:"phone" bson:"phone"`
		Code  string `json:"code" bson:"code"`
	}

	Item struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Name   string        `json:"name" bson:"name"`
		Image  string        `json:"image" bson:"image"`
		Price  float64       `json:"price" bson:"price"`
		Size   float64       `json:"size" bson:"size"`
		Desc   string        `json:"desc" bson:"desc"`
		Origin string        `json:"origin" bson:"origin"`
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
		Uid     string        `json:"uid" bson:"uid"`
		Name    string        `json:"name" bson:"name"`
		Mobile  string        `json:"mobile" bson:"mobile"`
		Region  string        `json:"region" bson:"region"`
		Address string        `json:"address" bson:"address"`
		Default bool          `json:"default" bson:"default"`
	}

	OrderItem struct {
		Name  string `json:"name" bson:"name"`
		Price string `json:"price" bson:"price"`
		Count string `json:"count" bson:"count"`
	}

	Order struct {
		Id      bson.ObjectId `json:"id" bson:"_id"`
		Uid     string        `json:"uid" bson:"uid"`
		Number  string        `json:"number" bson:"number"`
		Items   []*OrderItem  `json:"items" bson:"items"`
		Price   string        `json:"price" bson:"price"`
		Freight string        `json:"freight" bson:"freight"`
		Weight  string        `json:"weight" bson:"weight"`
		Address string        `json:"address" bson:"address"`
		Status  string        `json:"status" bson:"status"`
		Created int64         `json:"created" bson:"created"`
	}
)
