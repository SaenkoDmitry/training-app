package programs

import (
	"errors"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"time"

	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/programs"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/users"
)

type FindAllByUserUseCase struct {
	programsRepo programs.Repo
	usersRepo    users.Repo
}

func NewFindAllByUserUseCase(
	programsRepo programs.Repo,
	usersRepo users.Repo,
) *FindAllByUserUseCase {
	return &FindAllByUserUseCase{
		programsRepo: programsRepo,
		usersRepo:    usersRepo,
	}
}

func (uc *FindAllByUserUseCase) Name() string {
	return "Показать программы"
}

var (
	NoProgramsErr = errors.New("no training programs")
)

func (uc *FindAllByUserUseCase) Execute(chatID int64) (*dto.GetAllPrograms, error) {
	user, err := uc.usersRepo.GetByChatID(chatID)
	if err != nil {
		return nil, err
	}

	programObjs, err := uc.programsRepo.FindAll(user.ID)
	if err != nil {
		return nil, err
	}

	if len(programObjs) == 0 {
		return nil, NoProgramsErr
	}
	return &dto.GetAllPrograms{
		User:     user,
		Programs: mapToProgramDTO(programObjs),
	}, nil
}

func mapToProgramDTO(objs []models.WorkoutProgram) []*dto.ProgramDTO {
	result := make([]*dto.ProgramDTO, 0, len(objs))
	for _, obj := range objs {
		dayTypes := make([]*dto.WorkoutDayTypeDTO, 0, len(obj.DayTypes))
		for _, d := range obj.DayTypes {
			dayTypes = append(dayTypes, &dto.WorkoutDayTypeDTO{
				ID:               d.ID,
				WorkoutProgramID: d.WorkoutProgramID,
				Name:             d.Name,
				Preset:           d.Preset,
				CreatedAt:        "📅 " + d.CreatedAt.Add(time.Hour*3).Format("02.01.2006 15:04"),
			})
		}
		result = append(result, &dto.ProgramDTO{
			ID:        obj.ID,
			UserID:    obj.UserID,
			Name:      obj.Name,
			CreatedAt: obj.CreatedAt.Add(time.Hour * 3).Format("02.01.2006 15:04"),
			DayTypes:  dayTypes,
		})
	}
	return result
}
