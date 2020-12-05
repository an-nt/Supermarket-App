package Server

import (
	"SupermarketApp/API"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HttpsServer struct {
	http.Server
	Exec API.IApi
}
type IHttpsServer interface {
	StartServer()
}

func (sv *HttpsServer) StartServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome SSL server"))
	}).Methods("GET")

	router.HandleFunc("/v1/calculatemoney", sv.FormatCalMoney).Methods("POST")
	router.HandleFunc("/v1/stockchanges", sv.FormatChangeStock).Methods("POST")

	cert, err := tls.LoadX509KeyPair("Key/mysupermarket.crt", "Key/mysupermarket.key")
	if err != nil {
		panic(err)
	}

	sv.Handler = router
	sv.Addr = ":443"
	sv.ReadTimeout = 15 * time.Second
	sv.WriteTimeout = 15 * time.Second
	sv.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	log.Fatal(sv.ListenAndServeTLS("", ""))
}
