package repository

import (
	"github.com/rheadavin/hr-go-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// create
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// get all
func (r *UserRepository) FindAll(page, limit int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	query := r.db.Model(&models.User{})

	// Filter search
	if search != "" {
		query = query.Where(
			"name ILIKE ? OR email ILIKE ?",
			"%"+search+"%", "%"+search+"%",
		)
	}

	// Count total
	query.Count(&total)

	// Pagination
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&users).Error

	return users, total, err
}

// get by id
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err // gorm.ErrRecordNotFound jika tidak ada
	}
	return &user, nil
}

// get by email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err // gorm.ErrRecordNotFound jika tidak ada
	}
	return &user, nil
}

// update
func (r *UserRepository) Update(id uint, data map[string]interface{}) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
}

// delete
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
