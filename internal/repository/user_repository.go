package repository

import (
	"portfoliosite_v4_admin_auth_service/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

// CreateUser adds a new user to the database
func (repo *UserRepository) CreateUser(user *models.User) error {
    return repo.DB.Create(user).Error
}

// ListAllUsers retrieves all users from the database
func (repo *UserRepository) ListAllUsers() ([]models.User, error) {
    var users []models.User
    result := repo.DB.Find(&users)
    return users, result.Error
}

// ListUsersByRole retrieves users by specific role without sensitive data
func (repo *UserRepository) ListUsersByRole(role string) ([]models.User, error) {
    var users []models.User
    result := repo.DB.Where("role = ?", role).Find(&users)
    return users, result.Error
}

// GetUserByID retrieves a single user by ID, including sensitive data
func (repo *UserRepository) GetUserByID(id string) (*models.User, error) {
    var user models.User
    result := repo.DB.Where("id = ?", id).First(&user)
    return &user, result.Error
}

// UpdateUser updates an existing user
func (repo *UserRepository) UpdateUser(user *models.User) error {
    return repo.DB.Save(user).Error
}

// DeleteUser deletes a user by ID
func (repo *UserRepository) DeleteUser(id string) error {
    result := repo.DB.Delete(&models.User{}, "id = ?", id)
    return result.Error
}

// Retrieves a single user by email
func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := repo.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

