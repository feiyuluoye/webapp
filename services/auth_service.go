package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"tarot-app/models"
	"tarot-app/repositories"
)

type AuthService struct {
	UserRepository *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepository: repo}
}

// Authenticate 用户认证
func (s *AuthService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.UserRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	// 比较密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

// Register 用户注册
func (s *AuthService) Register(username, password string) (*models.User, error) {
	// 检查用户是否已存在
	existingUser, _ := s.UserRepository.FindByUsername(username)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// 保存用户到数据库
	err = s.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
