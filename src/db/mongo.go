package db

import (
	//"encoding/json"
	"errors"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	session *mgo.Session
}

func NewMongoDB(url string) *MongoDB {
	s, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	return &MongoDB{
		session: s,
	}
}

func (mg *MongoDB) GetItems(length int, lastid string) ([]Item, error) {
	items := []Item{}

	if lastid == "0" {
		mg.session.DB("").C("items").Find(nil).Limit(length).Sort("_id").All(&items)
	} else {
		if !bson.IsObjectIdHex(lastid) {
			return nil, errors.New("invalid id")
		}
		mg.session.DB("").C("items").Find(bson.M{"_id": bson.M{"$gt": bson.ObjectIdHex(lastid)}}).Limit(length).Sort("_id").All(&items)
	}

	return items, nil
}

func (mg *MongoDB) GetItem(id string) (*Item, error) {
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		log.Printf("[ERROR] invalid id %s", id)
		return nil, errors.New("invalid id")
	}

	oid := bson.ObjectIdHex(id)

	item := Item{}

	if err := mg.session.DB("").C("items").FindId(oid).One(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

// func (mg *MongoDB) GetOrder(id string) (*types.Order, error) {
// 	if !bson.IsObjectIdHex(id) {
// 		return nil, errors.New("invalid id")
// 	}
//
// 	order := types.Order{}
//
// 	if err := mg.session.DB().C("orders").FindId(bson.ObjectIdHex(id)).One(&order); err != nil {
// 		return nil, err
// 	}
//
// 	return order, nil
// }
//
// func (mg *MongoDB) GetOrders(length, lastid) (*types.Order, error) {
// 	if !bson.IsObjectIdHex(lastid) {
// 		return nil, errors.New("invalid id")
// 	}
//
// 	orders := []types.Order{}
// 	err := mg.session.DB().C("orders").Find(bson.M{"_id": bson.M{"$gt": bson.ObjectIdHex(lastid)}}).Limit(length).Sort("_id").All(&items)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return orders, nil
// }
//
// func (mg *MongoDB) NewOrder(body string) (string, error) {
// 	order := types.Order{}
//
// 	json.NewDecoder(body).Decode(&order)
//
// 	order.Id = bson.NewObjectId()
//
// 	if err := mg.session.DB().C("orders").Insert(order); err != nil {
// 		return nil, err
// 	}
//
// 	return json.Marshal(order)
// }
//
// func (mg *MongoDB) GetAddress(id string) (*types.Address, error) {
// 	if !bson.IsObjectIdHex(id) {
// 		return nil, errors.New("invalid id")
// 	}
//
// 	address := types.Address{}
//
// 	if err := mg.session.DB().C("address").FindId(bson.ObjectIdHex(id)).One(&address); err != nil {
// 		return nil, err
// 	}
//
// 	return address, nil
// }
//
// func (mg *MongoDB) NewAddress(body string) (string, error) {
// 	address := types.Address{}
//
// 	json.Decoder(body).Decode(&address)
//
// 	address.Id = bson.NewObjectId()
//
// 	if err := mg.session.DB().C("address").Insert(address); err != nil {
// 		return "", err
// 	}
//
// 	return json.Marshal(address)
// }
