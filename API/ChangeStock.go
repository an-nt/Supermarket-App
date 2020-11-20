package API

import (
	"SupermarketApp/Model"
)

func (a *Api) CheckRequestInfor(detail Model.StockHistory) error {
	if len(detail.ProductCode) != 3 {
		return Model.HaveError("Invalid product")
	}
	err := a.DbAccess.ChangeStock(detail)
	return err
}
