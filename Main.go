package main

import (
	"SupermarketApp/API"
	"SupermarketApp/Database"
	"SupermarketApp/Server"
	"fmt"
)

var ubuntuhost = "192.168.255.1"
var localhost = "localhost"
var containerhost = "172.18.0.2"

func main() {
	ms := Database.MSSQL{}
	db, result := ms.Config(containerhost, "sa", "khtn@2020", "1234", "Supermarket").Connect()
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
