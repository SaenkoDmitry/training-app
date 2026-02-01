package measurements

import (
	"time"

	"gorm.io/gorm"

	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
)

type Repo interface {
	Create(userID int64, shoulders, chest, hands, waist, buttocks, hips, calves, weight int) (*models.Measurement, error)
	Save(measurement *models.Measurement) error
	Get(measurementID int64) (models.Measurement, error)
	Delete(measurement *models.Measurement) error
	FindAll(userID int64) ([]models.Measurement, error)
	FindAllLimitOffset(userID int64, limit, offset int) ([]models.Measurement, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repoImpl{
		db: db,
	}
}

func (u *repoImpl) Create(userID int64, shoulders, chest, hands, waist, buttocks, hips, calves, weight int) (*models.Measurement, error) {
	newMeasurement := &models.Measurement{
		UserID:    userID,
		Shoulders: shoulders,
		Chest:     chest,
		Hands:     hands,
		Waist:     waist,
		Buttocks:  buttocks,
		Hips:      hips,
		Calves:    calves,
		Weight:    weight,
		CreatedAt: time.Now(),
	}
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&newMeasurement).Error
	})
	return newMeasurement, err
}

func (u *repoImpl) Delete(measurement *models.Measurement) error {
	return u.db.Delete(measurement).Error
}

func (u *repoImpl) Save(measurement *models.Measurement) error {
	return u.db.Save(measurement).Error
}

func (u *repoImpl) Get(measurementID int64) (measurement models.Measurement, err error) {
	tx := u.db.
		First(&measurement, measurementID)
	if tx.Error != nil {
		return models.Measurement{}, tx.Error
	}

	return measurement, nil
}

func (u *repoImpl) FindAll(userID int64) (measurements []models.Measurement, err error) {
	tx := u.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&measurements)

	if tx.Error != nil {
		return []models.Measurement{}, tx.Error
	}

	return measurements, nil
}

func (u *repoImpl) FindAllLimitOffset(userID int64, limit, offset int) (measurements []models.Measurement, err error) {
	tx := u.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&measurements)

	if tx.Error != nil {
		return []models.Measurement{}, tx.Error
	}

	return measurements, nil
}
