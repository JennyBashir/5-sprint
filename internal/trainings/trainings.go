package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	energy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	//"3456,Ходьба,3h00m"
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("length of data is not 3")
	}
	t.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("error in convert string to int")
	}
	if t.Steps <= 0 {
		return fmt.Errorf("wrong number of steps")
	}

	t.Duration, err = time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("error in convert string to duration")
	}
	if t.Duration <= 0 {
		return fmt.Errorf("invalid duration value")
	}

	t.TrainingType = slice[1]
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := energy.Distance(t.Steps, t.Height)
	speed := energy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		kcal, err := energy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error in calories calc")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, kcal), nil

	case "Бег":
		kcal, err := energy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error in calories calc")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, kcal), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
