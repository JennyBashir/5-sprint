package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	energy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	//"678,0h50m"
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("length of data is not 2")
	}
	ds.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("error in convert string to int")
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("wrong number of steps")
	}

	ds.Duration, err = time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("error in convert string to duration")
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("invalid duration value")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := energy.Distance(ds.Steps, ds.Height)
	kcal, err := energy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error in calories calc")
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, kcal), nil
}
