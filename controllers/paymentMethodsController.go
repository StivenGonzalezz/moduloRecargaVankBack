package controllers

import (
	"encoding/json"
	"modulo_recarga/services"
	"net/http"
)

type PaymentMethodsController struct{}

func (p *PaymentMethodsController) GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Cuerpo de solicitud inválido", http.StatusBadRequest)
		return
	}

	id, ok := requestBody["id"].(string)
	if !ok {
		http.Error(w, "ID no proporcionado o es inválido", http.StatusBadRequest)
		return
	}

	// Llama a getPaymentMethods del servicio
	result, err := services.GetPaymentMethods(id)
	if err != nil {
		http.Error(w, "Error al obtener métodos de pago", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "ok",
		"result":  result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
