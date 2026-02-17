package sets

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"gorm.io/gorm"
)

type Repo interface {
	Get(id int64) (*models.Set, error)
	Delete(id int64) error
	DeleteAllBy(exerciseID int64) error
	Save(set *models.Set) error
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repoImpl{
		db: db,
	}
}

func (u *repoImpl) Get(id int64) (*models.Set, error) {
	var set models.Set
	u.db.Preload("Exercise").Preload("Exercise.ExerciseType").First(&set, id)
	return &set, nil
}

func (u *repoImpl) Delete(id int64) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&models.Set{}).Error
	})
}

func (u *repoImpl) DeleteAllBy(exerciseID int64) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("exercise_id = ?", exerciseID).Delete(&models.Set{}).Error
	})
}

func (u *repoImpl) Save(set *models.Set) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Save(&set).Error
	})
}
