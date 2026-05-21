package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию

	// Преобразуем строку в слайс строк и проверяем длину
	sliceStr := strings.Split(data, ",")
	if len(sliceStr) != 2 {
		return 0, 0, fmt.Errorf("invalid length, got %d, expected 2", len(sliceStr))
	}

	steps, err := strconv.Atoi(sliceStr[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error conversion steps: %w", err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("number of steps must be positive and greater than 0")
	}

	period, err := time.ParseDuration(sliceStr[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing duration: %w", err)
	}

	if period <= 0 {
		return 0, 0, fmt.Errorf("duration must be positive and greater than 0")
	}

	return steps, period, nil
}

// DayActionInfo с помощью parsePackage парсит строку с данными.
// Вычисляет дистанцию в километрах и количество потраченных калорий
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	// Получаем данные о кол-ве шагов и продолжительности
	steps, period, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * stepLength
	distanceInKm := distance / mInKm

	burnedCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, period)
	if err != nil {
		log.Printf("error calculating burned calories: %v", err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceInKm, burnedCalories)
}
