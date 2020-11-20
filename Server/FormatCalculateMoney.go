package Server

import (
	"encoding/json"
	"net/http"
)

type CalMoneyInput struct {
	Code string `json: "code"`
	Num  uint   `json: "quantity"`
}

type CalMoneyOutput struct {
	Vnd uint    `json: "VNDmoney"`
	Usd float32 `json: "USDmoney`
}

func (sv *HttpsServer) FormatCalMoney(w http.ResponseWriter, r *http.Request) {
	var in CalMoneyInput
	var out CalMoneyOutput

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		resp := ResponseFormat{
			Message: "Invalid type",
			Detail:  nil,
		}
		resp.FormatResponse(w, http.StatusBadRequest)
		return
	}

	rate, err := sv.Exec.GetExchangeRate()
	if err != nil {
		resp := ResponseFormat{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResponse(w, http.StatusBadRequest)
		return
	}

	price, err := sv.Exec.GetPrice(in.Code)
	if err != nil {
		resp := ResponseFormat{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResponse(w, http.StatusBadRequest)
		return
	}

	out.Vnd = in.Num * price
	out.Usd = float32(out.Vnd) / float32(rate)

	resp := ResponseFormat{
		Message: "Success",
		Detail:  out,
	}
	resp.FormatResponse(w, http.StatusOK)
}
