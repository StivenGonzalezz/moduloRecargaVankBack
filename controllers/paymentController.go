package controllers

import (
    "encoding/json"
    "net/http"
    "modulo_recarga/services"
)

type PaymentController struct{}

func (p *PaymentController) PostPayment(w http.ResponseWriter, r *http.Request) {
    // Decodifica el cuerpo de la solicitud JSON en un mapa o struct
    var requestBody map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Cuerpo de solicitud inválido", http.StatusBadRequest)
        return
    }

    requestBodyJSON, err := json.Marshal(requestBody)
    if err != nil {
        http.Error(w, "Error al convertir la solicitud en JSON", http.StatusInternalServerError)
        return
    }

    // Llama a la función postPayment del servicio de pagos
    result, err := services.GetPaymentMethods(string(requestBodyJSON))
    if err != nil {
        http.Error(w, "Error en el procesamiento del pago", http.StatusInternalServerError)
        return
    }

    // Crea la respuesta
    response := map[string]interface{}{
        "message": "ok",
        "result":  result,
    }

    // Codifica la respuesta como JSON y la envía al cliente
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

