package API

import (
	"SupermarketApp/Database"
	"SupermarketApp/Model"
)

type Api struct {
	DbAccess Database.IDatabaseAccess
}

type IApi interface {
	CheckRequestInfor(detail Model.StockHistory) error
	GetExchangeRate() (uint, error)
	GetPrice(productCode string) (uint, error)
}
