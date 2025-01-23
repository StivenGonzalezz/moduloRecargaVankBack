package usescases

import (
	"fmt"
	"modulo_recarga/infrastructure/db"
	"modulo_recarga/infrastructure/db/models"

	"github.com/mercadopago/sdk-go/pkg/payment"
)

func UpdateTransferStatus(response *payment.Response) {

	var payment models.Payment

	if err := db.DB.Unscoped().First(&payment, "id = ?", response.ID).Error; err != nil {
		fmt.Printf("Error al buscar el registro con ID %v: %v\n", response.ID, err)
		return
	}

	if err := db.DB.Unscoped().Model(&payment).Updates(models.Payment{Status:response.Status}).Error; err != nil {
		fmt.Printf("Error al actualizar el registro con ID %v: %v\n", response.ID, err)
		return
	}

	fmt.Printf("El registro con ID %v fue actualizado a: %v\n", response.ID, response.Status)
}