package Server

import (
	"SupermarketApp/Model"
	"encoding/json"
	"net/http"
)

func (sv *HttpsServer) FormatChangeStock(w http.ResponseWriter, r *http.Request) {
	var detail Model.StockHistory
	err := json.NewDecoder(r.Body).Decode(&detail)
	if err != nil {
		resp := ResponseFormat{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResponse(w, http.StatusBadRequest)
		return
	}

	err = sv.Exec.CheckRequestInfor(detail)
	if err != nil {
		resp := ResponseFormat{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResponse(w, http.StatusInternalServerError)
		return
	}

	resp := ResponseFormat{
		Message: "Success",
		Detail:  nil,
	}
	resp.FormatResponse(w, http.StatusOK)
}
