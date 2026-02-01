package measurements

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/measurements"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/users"
	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
)

type FindAllByUserUseCase struct {
	measurementsRepo measurements.Repo
	usersRepo        users.Repo
}

func NewFindAllByUserUseCase(
	measurementsRepo measurements.Repo,
	usersRepo users.Repo,
) *FindAllByUserUseCase {
	return &FindAllByUserUseCase{
		measurementsRepo: measurementsRepo,
		usersRepo:        usersRepo,
	}
}

func (uc *FindAllByUserUseCase) Name() string {
	return "Показать измерения"
}

func (uc *FindAllByUserUseCase) Execute(chatID int64, limit, offset int) ([]dto.Measurement, error) {
	user, err := uc.usersRepo.GetByChatID(chatID)
	if err != nil {
		return nil, err
	}

	measurementObjs, err := uc.measurementsRepo.FindAllLimitOffset(user.ID, limit, offset)
	if err != nil {
		return nil, err
	}

	result := make([]dto.Measurement, 0, len(measurementObjs))
	for _, m := range measurementObjs {
		result = append(result, dto.Measurement{
			CreatedAt: utils.FormatDate(m.CreatedAt),
			Shoulders: utils.FormatCentimeters(m.Shoulders),
			Chest:     utils.FormatCentimeters(m.Chest),
			Hands:     utils.FormatCentimeters(m.Hands),
			Waist:     utils.FormatCentimeters(m.Waist),
			Buttocks:  utils.FormatCentimeters(m.Buttocks),
			Hips:      utils.FormatCentimeters(m.Hips),
			Calves:    utils.FormatCentimeters(m.Calves),
			Weight:    utils.FormatKilograms(m.Weight),
		})
	}

	return result, nil
}
