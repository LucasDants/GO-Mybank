package models

import (
	"errors"
	"time"
	"strings"
	"myBank/src/security"
)

type Account struct {
	ID         uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Password  string    `json:"password,omitempty"`
	Currency string 	`json:"currency,omitempty"`
	Amount      float64    `json:"amount"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (account *Account) Prepare() error {
	if err := account.validate(); err != nil {
		return err
	}

	if err := account.format(); err != nil {
		return err
	}

	return nil
}

func (account *Account) validate() error {
	if account.Name == "" {
		return errors.New("nome é obrigatório e não pode estar em branco")
	}

	if account.Password == "" {
		return errors.New("senha é obrigatória e não pode estar em branco")
	}

	if account.Currency == "" {
		return errors.New("moeda é obrigatório e não pode estar em branco")
	}

	if account.Currency != "BRL" && account.Currency != "GBP" && account.Currency != "USD" && account.Currency != "EUR" {
		return errors.New("não fornecemos o tipo de moeda selecionado.")
	}


	return nil
}

func (account *Account) format() error {
	account.Name = strings.TrimSpace(account.Name)
	account.Currency = strings.TrimSpace(account.Currency)

		hashPassword, err := security.Hash(account.Password)
		if err != nil {
			return err
		}

		account.Password = string(hashPassword)
	return nil
}