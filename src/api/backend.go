package api

import (
	"github.com/wlshcy/rabbit/src/db"
)

type Backend interface {
	GetItem(string) (*db.Item, error)
	GetItems(int, string) ([]db.Item, error)
	GetOnSale(string) (*db.OnSale, error)
	GetOnSales() ([]db.OnSale, error)
	GetAddresses() ([]db.Address, error)
	GetAddress(string) (*db.Address, error)
	NewAddress(*db.Address) error
}
