package API

func (a *Api) GetPrice(productCode string) (uint, error) {
	price, err := a.DbAccess.GetProductPrice(productCode)
	if err != nil {
		return 0, err
	}

	return price, nil
}
