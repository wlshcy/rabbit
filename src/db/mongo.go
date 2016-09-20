package db

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/wlshcy/rabbit/src/auth"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	session *mgo.Session
}

const (
	SMS_TTL = 60
)

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

func (mg *MongoDB) GetOnSales() ([]OnSale, error) {
	onsales := []OnSale{}

	if err := mg.session.DB("").C("onsales").Find(nil).Sort("_id").All(&onsales); err != nil {
		log.Printf("[ERROR] Retrive on-sales failed: %s", err)
		return nil, err
	}

	return onsales, nil
}

func (mg *MongoDB) GetOnSale(id string) (*OnSale, error) {
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		log.Printf("[ERROR] invalid id %s", id)
		return nil, errors.New("invalid id")
	}

	oid := bson.ObjectIdHex(id)

	onsale := OnSale{}

	if err := mg.session.DB("").C("onsales").FindId(oid).One(&onsale); err != nil {
		log.Printf("[ERROR]  Find %s from mongo failed: %s", id, err)
		return nil, err
	}

	return &onsale, nil
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
func (mg *MongoDB) GetAddresses(uid string) ([]Address, error) {
	addresses := []Address{}

	if err := mg.session.DB("").C("addresses").Find(bson.M{"uid": uid}).Sort("_id").All(&addresses); err != nil {
		log.Printf("[ERROR] Retrive addresses failed: %s", err)
		return nil, err
	}

	return addresses, nil
}

func (mg *MongoDB) GetAddress(id string) (*Address, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid id")
	}

	address := Address{}

	if err := mg.session.DB("").C("addresses").FindId(bson.ObjectIdHex(id)).One(&address); err != nil {
		return nil, err
	}

	return &address, nil
}

func (mg *MongoDB) NewAddress(address *Address) error {
	address.Id = bson.NewObjectId()

	if err := mg.session.DB("").C("addresses").Insert(address); err != nil {
		return err
	}
	return nil
}

func (mg *MongoDB) UpdateAddress(id string, body map[string]interface{}) error {
	if !bson.IsObjectIdHex(id) {
		log.Printf("[ERROR] Invalid id %s", id)
		return errors.New("invalid id")
	}

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": body}
	if err := mg.session.DB("").C("addresses").Update(query, change); err != nil {
		log.Printf("[ERROR] Update address %s failed: %s", id, err)
		return err
	}
	// How to use Apply?
	// change := mgo.Change{
	// 	Update:    bson.M{"$set": body},
	// 	ReturnNew: true,
	// }
	// mg.session.DB("").C("addresses").Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &address)

	return nil
}

func (mg *MongoDB) DeleteAddress(id string) error {
	if !bson.IsObjectIdHex(id) {
		log.Printf("[ERROR] Invalid id %s", id)
		return errors.New("invalid id")
	}

	if err := mg.session.DB("").C("addresses").Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		log.Printf("[ERROR] Delete address %s failed: %s", id, err)
		return err
	}

	return nil
}

func (mg *MongoDB) Login(credential *Credential) (string, error) {
	authBackend := auth.NewJWTAuthenticationBackend()

	token, err := authBackend.GenerateToken(credential.Phone)
	if err != nil {
		return "", err
	}

	n, _ := mg.session.DB("").C("users").Find(bson.M{"phone": credential.Phone}).Count()
	if n == 0 {
		user := User{
			Phone: credential.Phone,
		}

		if err := mg.session.DB("").C("users").Insert(&user); err != nil {
			log.Printf("[ERROR] Insert user failed: %s", err)
			return "", err
		}
	}

	return token, nil
}

func (mg *MongoDB) SendSMS(to string) error {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "0123456789"
	result := make([]byte, 6)
	for i := 0; i < 6; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	sms := string(result)

	if err := mg.session.DB("").C("sms").Insert(bson.M{
		"sms": sms,
	}); err != nil {
		log.Printf("[ERROR] Insert sms code failed: %s", err)
		return err
	}
	//sessionTTL := mgo.Index{
	//	Key:         []string{"created"},
	//	Unique:      false,
	//	DropDups:    false,
	//	Background:  true,
	//	ExpireAfter: time.Duration(60) * time.Second,
	//}
	//mg.session.DB("").C("token").EnsureIndex(sessionTTL)
	go func() {
		//ticker := time.NewTicker(time.Second * time.Duration(60))
		//<-ticker.C
		time.Sleep(time.Second * time.Duration(SMS_TTL))
		mg.session.DB("").C("sms").Remove(bson.M{"sms": sms})
		//ticker.Stop()
	}()

	return nil
}
