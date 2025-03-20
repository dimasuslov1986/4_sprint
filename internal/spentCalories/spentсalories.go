package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	sliseStrings := strings.Split(data, ",")
	if len(sliseStrings) != 3 {
		return 0, "0", 0, fmt.Errorf("Ошибка")
	}

	var steps int
	var err error
	steps, err = strconv.Atoi(sliseStrings[0])
	if steps <= 0 || err != nil {
		return 0, "0", 0, fmt.Errorf("Ошибка")
	}

	duration, err := time.ParseDuration(sliseStrings[2])
	if err != nil {
		return 0, "0", 0, fmt.Errorf("Ошибка")
	}
	return steps, sliseStrings[1], duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	dist := float64(steps) * lenStep / mInKm
	return dist
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0 {
		return 0
	}
	dist := distance(steps)
	meanSp := dist / duration.Hours()
	return meanSp
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, active, duration, err := parseTraining(data)
	if err != nil {
		return "ошибка"
	}
	switch active {
	case "Ходьба":
		dist := distance(steps)
		meanSp := meanSpeed(steps, duration)
		ccal := WalkingSpentCalories(steps, weight, height, duration)
		info := fmt.Sprintf(`Тип тренировки: %s
		Длительность: %.2f ч.
		Дистанция: %.2f км.
		Скорость: %.2f км/ч
		Сожгли калорий: %.2f`, active, duration.Hours(), dist, meanSp, ccal)
		return info
	case "Бег":
		dist := distance(steps)
		meanSp := meanSpeed(steps, duration)
		ccal := RunningSpentCalories(steps, weight, duration)
		info := fmt.Sprintf(`Тип тренировки: %s
		Длительность: %.2f ч.
		Дистанция: %.2f км.
		Скорость: %.2f км/ч
		Сожгли калорий: %.2f`, active, duration.Hours(), dist, meanSp, ccal)
		return info
	default:
		return "неизвестный тип тренировки"
	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSp := meanSpeed(steps, duration)
	ccal := ((runningCaloriesMeanSpeedMultiplier * meanSp) - runningCaloriesMeanSpeedShift) * weight
	return ccal

}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSp := meanSpeed(steps, duration)
	ccal := ((walkingCaloriesWeightMultiplier * weight) + (meanSp*meanSp/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
	return ccal
}
