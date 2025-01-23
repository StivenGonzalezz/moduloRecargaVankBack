package models

import "gorm.io/gorm"

type Payment struct {
	ID                        uint    `gorm:"primaryKey"`
	ExternalReference         string  `gorm:"not null"`
	TransactionAmount         float64 `gorm:"not null"`
	Description               string
	PaymentMethodID           string
	PayerEmail                string
	PayerIdentificationType   string
	PayerIdentificationNumber string
	ZipCode                   string
	StreetName                string
	StreetNumber              string
	Neighborhood              string
	City                      string
	FederalUnit               string
	PhoneAreaCode             string
	PhoneNumber               string
	FinancialInstitution      string
	Status                    string
	CreatedAt                 *gorm.DeletedAt
	UpdatedAt                 *gorm.DeletedAt
}

// Write implements io.Writer.
func (Payment) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}
