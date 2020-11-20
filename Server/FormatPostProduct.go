package Server

import (
	"net/http"
)

type PostProductFormatter struct {
}

type IPostProductFormatter interface {
	FormatPostProduct(w http.ResponseWriter, r *http.Request)
}

func (f *PostProductFormatter) FormatPostProduct(w http.ResponseWriter, r *http.Request) {

}
