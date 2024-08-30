package repositories

import (
	"errors"
	initial "tarot-app/initial"
	"tarot-app/models"
)

type TarotSpreadRepository struct {
	spreads []models.TarotSpread
	nextID  int
}

// NewTarotSpreadRepository initializes a new repository with sample data
func NewTarotSpreadRepository() *TarotSpreadRepository {

	return &TarotSpreadRepository{
		spreads: initial.InitializeTarotSpreads(),
	}
}

// Create adds a new TarotSpread to the repository
func (r *TarotSpreadRepository) Create(spread models.TarotSpread) models.TarotSpread {

	r.spreads = append(r.spreads, spread)
	return spread
}

// Read returns a TarotSpread by its ID
func (r *TarotSpreadRepository) Read(id int) (*models.TarotSpread, error) {
	for _, spread := range r.spreads {
		if spread.ID == uint(id) {
			return &spread, nil
		}
	}
	return nil, errors.New("spread not found")
}

// Update modifies an existing TarotSpread
func (r *TarotSpreadRepository) Update(id int, updatedSpread models.TarotSpread) (*models.TarotSpread, error) {
	for i, spread := range r.spreads {
		if spread.ID == uint(id) {
			updatedSpread.ID = uint(id)
			r.spreads[i] = updatedSpread
			return &updatedSpread, nil
		}
	}
	return nil, errors.New("spread not found")
}

// Delete removes a TarotSpread by its ID
func (r *TarotSpreadRepository) Delete(id int) error {
	for i, spread := range r.spreads {
		if spread.ID == uint(id) {
			r.spreads = append(r.spreads[:i], r.spreads[i+1:]...)
			return nil
		}
	}
	return errors.New("spread not found")
}

// List returns all TarotSpreads
func (r *TarotSpreadRepository) List() []models.TarotSpread {
	return r.spreads
}
