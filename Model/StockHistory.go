package Model

type StockHistory struct {
	ProductCode string //char(3)
	InWarehouse int64  //int
	Day         string //yyyymmdd
	Time        string //hh:mm:ss.mmm
	Change      int64  //not zero
	DoneBy      int64  //int
}
