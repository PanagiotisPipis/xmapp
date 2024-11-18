package models

import (
	"github.com/google/uuid"
)

type Company struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string    `gorm:"size:15;unique;not null" json:"name"`
	Description    string    `gorm:"size:3000" json:"description"`
	AmountOfEmployees int     `gorm:"not null" json:"amount_of_employees"`
	Registered     bool      `gorm:"not null" json:"registered"`
	Type           string    `gorm:"not null" json:"type" enums:"Corporations,NonProfit,Cooperative,Sole Proprietorship"`
}
