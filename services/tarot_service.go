package services

import (
	"tarot-app/models"
	"tarot-app/repositories"
)

type TarotService struct {
	Repository *repositories.TarotRepository
}

func NewTarotService(repo *repositories.TarotRepository) *TarotService {
	return &TarotService{Repository: repo}
}

func (s *TarotService) GetAllCards() ([]models.TarotCard, error) {
	return s.Repository.GetAll()
}

func (s *TarotService) GetCardByID(id uint) (*models.TarotCard, error) {
	return s.Repository.GetByID(id)
}
