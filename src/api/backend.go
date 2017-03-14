package api

import (
	"github.com/wlshcy/rabbit/src/db"
)

type Backend interface {
	Login(*db.Credential) (string, error)

	CreateItem(*db.Item) error
	GetItem(string) (*db.Item, error)
	GetItems(int, string) ([]db.Item, error)
	UpdateItem(*db.Item) error
	DeleteItem(string) error

	GetOnSale(string) (*db.OnSale, error)
	GetOnSales() ([]db.OnSale, error)

	GetAddresses(string) ([]db.Address, error)
	GetAddress(string) (*db.Address, error)
	UpdateAddress(string, map[string]interface{}) error
	DeleteAddress(string) error
	DefaultAddress(string, string) error
	NewAddress(*db.Address) error

	SendSMS(string) error

	GetOrder(string) (*db.Order, error)
	NewOrder(*db.Order) error
}
