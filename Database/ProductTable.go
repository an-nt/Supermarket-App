package Database

import (
	"SupermarketApp/Model"
	"errors"
	"fmt"
)

func (ms *MSSQL) GetProductPrice(code string) (uint, error) {
	var pro Model.Product
	var proIdx []Model.Product

	query := fmt.Sprintf("SELECT * FROM dbo.Product WHERE Code = '%s';", code)

	rows, err := ms.Db.Query(query)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&pro.Code, &pro.Name, &pro.Type, &pro.Price)
		if err != nil {
			fmt.Println(err) //maybe have a minor error
		}
		proIdx = append(proIdx, pro)
	}

	switch len(proIdx) {
	case 0:
		return 0, errors.New("Product invalid")
	case 1:
		return pro.Price, nil
	default:
		return 0, errors.New("Product invalid")
	}
}
