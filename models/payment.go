package models

import (
	"database/sql"
)

type PaymentService struct {
	DB *sql.DB
}

func (ps PaymentService) FulfillOrder(payload []byte, signatureHeader string) (bool, error) {

	return false, nil
}
