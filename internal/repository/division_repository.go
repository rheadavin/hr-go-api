package repository

import (
	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/models"
	"gorm.io/gorm"
)

type DivisionRepository struct {
	db *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) *DivisionRepository {
	return &DivisionRepository{db: db}
}

func (r *DivisionRepository) Create(division *models.Division) error {
	return r.db.Create(division).Error
}

func (r *DivisionRepository) FindAll(page, limit int, search string) ([]dto.DivisionResponse, int64, error) {
	var divisions []dto.DivisionResponse
	var total int64
	query := r.db.Model(&models.Division{}).Select("id", "name").Where("is_active = ?", true)

	// Filter search
	if search != "" {
		query = query.Where(
			"name ILIKE ?",
			"%"+search+"%",
		)
	}

	// Count total
	query.Count(&total)

	// Pagination
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).
		Order("created_at DESC").
		Scan(&divisions).Error

	return divisions, total, err
}

func (r *DivisionRepository) FindByID(id uint) (*models.Division, error) {
	var division models.Division
	err := r.db.Model(&models.Division{}).Where("id = ? AND is_active = ?", id, true).First(&division).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &division, nil
}

func (r *DivisionRepository) Update(id uint, data *models.Division) error {
	return r.db.Model(&models.Division{}).Where("id = ? AND is_active = ?", id, true).Updates(data).Error
}

func (r *DivisionRepository) Delete(id uint) error {
	return r.db.Model(&models.Division{}).Where("id = ? AND is_active = ?", id, true).Updates(map[string]interface{}{"is_active": false}).Error
}
