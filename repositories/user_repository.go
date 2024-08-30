package repositories

import (
	"gorm.io/gorm"
	"tarot-app/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// FindByUsername 根据用户名查找用户
func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Save 保存用户到数据库
func (repo *UserRepository) Save(user *models.User) error {
	return repo.DB.Create(user).Error
}
