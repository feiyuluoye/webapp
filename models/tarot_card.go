package models

import "gorm.io/gorm"

type TarotCard struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Arcana      string `json:"arcana"`
	Suit        string `json:"suit"`
	Number      int    `json:"number"`
}
