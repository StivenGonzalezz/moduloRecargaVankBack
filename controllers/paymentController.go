package controllers

import (
	"encoding/json"
	"io/ioutil"
	"modulo_recarga/services"
	"net/http"

	"gorm.io/gorm"
)

type PaymentController struct{
	db *gorm.DB
}

func (p *PaymentController) PostPayment(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Solicitud sin cuerpo", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	var requestBody map[string]interface{}
	if err := json.Unmarshal(body, &requestBody); err != nil {
		http.Error(w, "Cuerpo de solicitud inválido", http.StatusBadRequest)
		return
	}

	result, err := services.PostPayment(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
