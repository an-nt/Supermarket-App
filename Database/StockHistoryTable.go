package Database

import (
	"SupermarketApp/Model"
	"fmt"
)

func (ms *MSSQL) ChangeStock(detail Model.StockHistory) error {
	query := fmt.Sprintf(`
	INSERT [dbo].[StockHistory] ([ProductCode],[InWarehouse],[Day],[Time],[Change],[DoneBy])
	VALUES ('%s',%d,getdate(),Getdate(),%d,%d)
	`, detail.ProductCode, detail.InWarehouse, detail.Change, detail.DoneBy)
	_, err := ms.Db.Exec(query)
	return err
}
