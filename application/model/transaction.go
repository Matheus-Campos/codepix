package model

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID           string  `json:"id" validate:"required,uuidv4"`
	AccountID    string  `json:"accountId" validate:"required,uuidv4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error"`
}

func (t *Transaction) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		_ = fmt.Errorf("error during Transaction validation: %s", err.Error())
		return err
	}
	return nil
}

func (t *Transaction) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	return t.isValid()
}

func (t *Transaction) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	return json.Marshal(t)
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
