package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	ErrInput = errors.New("input data error")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("wrong number of steps: %w", ErrInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration value: %w", ErrInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("wrong value of weight: %w", ErrInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("wrong value of height: %w", ErrInput)
	}
	speed := MeanSpeed(steps, height, duration)
	calories := ((weight * speed * duration.Minutes()) / float64(minInH)) * walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("wrong number of steps: %w", ErrInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration value: %w", ErrInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("wrong value of weight: %w", ErrInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("wrong value of height: %w", ErrInput)
	}
	speed := MeanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / float64(minInH)
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	dist := Distance(steps, height)
	averageSpeed := dist / duration.Hours()
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepsLength := height * stepLengthCoefficient
	distance := (float64(steps) * stepsLength) / float64(mInKm)
	return distance
}
