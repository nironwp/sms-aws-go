package routes

import (
	"encoding/json"
	"net/http"
	snsmananger "sms-aws-go/cmd/utils"
)

func SendSMSHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		PhoneNumber string `json:"phone"`
		Message     string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send SMS using sendSMS function
	err = snsmananger.SendSMS(reqBody.Message, reqBody.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	resBody := struct {
		Message string `json:"message"`
	}{
		Message: "SMS ENVIADO!",
	}
	resBodyJSON, _ := json.Marshal(resBody)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resBodyJSON)
}
