package Database

import (
	"SupermarketApp/Model"
	"database/sql"
)

type Database struct {
	Db *sql.DB
}

type IDatabaseAccess interface {
	Connect() (*sql.DB, string)
	GetProductPrice(code string) (uint, error)
	ChangeStock(detail Model.StockHistory) error
}
