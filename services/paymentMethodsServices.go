package services

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "os"
)

func GetPaymentMethods(id string) ([]interface{}, error) {
    url := "https://api.mercadopago.com/v1/payment_methods"

    // Crear la solicitud HTTP
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("error al crear la solicitud: %v", err)
    }

    // Agregar encabezados
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ACCESS_TOKEN")))

    // Realizar la solicitud
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error al realizar la solicitud: %v", err)
    }
    defer resp.Body.Close()

    // Comprobar el estado de la respuesta
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error HTTP! estado: %d", resp.StatusCode)
    }

    // Decodificar la respuesta JSON
    var data []map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, fmt.Errorf("error al decodificar JSON: %v", err)
    }

    // Filtrar el método de pago con el ID especificado
    var dataPSE map[string]interface{}
    for _, method := range data {
        if method["id"] == id {
            dataPSE = method
            break
        }
    }

    // Comprobar si el método fue encontrado y devolver instituciones financieras
    if dataPSE == nil {
        return nil, errors.New("método de pago no encontrado")
    }

    financialInstitutions, ok := dataPSE["financial_institutions"].([]interface{})
    if !ok {
        return nil, errors.New("no se encontraron instituciones financieras")
    }

    return financialInstitutions, nil
}
