package models

import (
	"errors"
	"time"
)

type Operation struct {
	ID uint64 `json:"id,omitempty"`
	Origin         uint64    `json:"origin,omitempty"`
	Destiny      uint64    `json:"destiny,omitempty"`
	ExchangeCurrency  float64    `json:"exchangeCurrency,omitempty"`
	CurrencyTransformed  float64    `json:"currencyTransformed,omitempty"`
	OriginCurrency string  `json:"originCurrency,omitempty"`
	DestinyCurrency string  `json:"destinyCurrency,omitempty"`
	Type string `json:"type,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (operation *Operation) Prepare() error {
	if err := operation.validate(); err != nil {
		return err
	}

	return nil
}

func (operation *Operation) validate() error {
	if operation.ExchangeCurrency == 0 {
		return errors.New("não é possível realizar uma operação com valor nulo")
	}

	return nil
}
