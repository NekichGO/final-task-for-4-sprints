package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию

	sliceStr := strings.Split(data, ",")

	if len(sliceStr) != 3 {
		return 0, "", 0, fmt.Errorf("invalid length slice: got %d, expected 3", len(sliceStr))
	}

	steps, err := strconv.Atoi(sliceStr[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("error conversion steps: %w", err)
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("number of steps must be positive and greater than 0")
	}

	period, err := time.ParseDuration(sliceStr[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("error parse duration: %w", err)
	}

	if period <= 0 {
		return 0, "", 0, fmt.Errorf("error parsing: duration must be positive and greater than 0")
	}

	typeOfTraining := sliceStr[1]

	return steps, typeOfTraining, period, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	stepsLength := height * stepLengthCoefficient

	return (float64(steps) * stepsLength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if duration <= 0 {
		return 0
	}

	way := distance(steps, height)

	avgSpeed := way / duration.Hours()

	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	// Получаем значение из строки
	steps, typeOfTraining, period, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Проверяем вид тренировки и производим расчеты для каждого вида
	trainingDistance := distance(steps, height)
	avgSpeed := meanSpeed(steps, height, period)
	switch typeOfTraining {
	case "Ходьба":
		walkCalories, err := WalkingSpentCalories(steps, weight, height, period)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfTraining, period.Hours(), trainingDistance, avgSpeed, walkCalories), nil

	case "Бег":
		runCalories, err := RunningSpentCalories(steps, weight, height, period)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfTraining, period.Hours(), trainingDistance, avgSpeed, runCalories), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", typeOfTraining)
	}
}

// RunningSpentCalories производит расчет затраченных калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// Проверяем входные параметры на корректность
	if steps <= 0 {
		return 0, fmt.Errorf("number of steps must be greater than 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than 0")
	}

	avgSpeedRun := meanSpeed(steps, height, duration)

	caloriesAtRunning := (weight * avgSpeedRun * duration.Minutes()) / minInH

	return caloriesAtRunning, nil
}

// WalkingSpentCalories производит расчет затраченных калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// Проверяем входные параметры на корректность
	if steps <= 0 {
		return 0, fmt.Errorf("numbers of steps must be greater than 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than 0")
	}

	avgSpeedWalk := meanSpeed(steps, height, duration)

	caloriesAtWalking := (weight * avgSpeedWalk * duration.Minutes()) / minInH

	return caloriesAtWalking * walkingCaloriesCoefficient, nil

}
