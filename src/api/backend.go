package api

import (
	"github.com/wlshcy/rabbit/src/db"
)

type Backend interface {
	Login(*db.Credential) (string, error)
	GetItem(string) (*db.Item, error)
	GetItems(int, string) ([]db.Item, error)
	GetOnSale(string) (*db.OnSale, error)
	GetOnSales() ([]db.OnSale, error)
	GetAddresses() ([]db.Address, error)
	GetAddress(string) (*db.Address, error)
	UpdateAddress(string, map[string]interface{}) error
	DeleteAddress(string) error
	NewAddress(*db.Address) error
	SendSMS(string) error
}
