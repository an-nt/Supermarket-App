package main

import (
	"SupermarketApp/API"
	"SupermarketApp/Database"
	"SupermarketApp/Server"
	"fmt"
)

var ubuntuhost = "192.168.255.1"
var localhost = "127.0.0.1"

func main() {
	ms := Database.MSSQL{}
	db, result := ms.Config(ubuntuhost, "sa", "khtn@2020", "1433", "Supermarket").Connect()
	fmt.Println(result)

	srv := Server.HttpsServer{
		Exec: &API.Api{
			DbAccess: &Database.MSSQL{
				Database: Database.Database{Db: db},
			},
		},
	}
	srv.StartServer()
}
