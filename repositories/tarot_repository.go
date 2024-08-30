package repositories

import (
	"gorm.io/gorm"
	"tarot-app/models"
)

type TarotRepository struct {
	DB *gorm.DB
}

func NewTarotRepository(db *gorm.DB) *TarotRepository {
	return &TarotRepository{DB: db}
}

func (repo *TarotRepository) GetAll() ([]models.TarotCard, error) {
	var cards []models.TarotCard
	err := repo.DB.Find(&cards).Error
	return cards, err
}

func (repo *TarotRepository) GetByID(id uint) (*models.TarotCard, error) {
	var card models.TarotCard
	err := repo.DB.First(&card, id).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}
