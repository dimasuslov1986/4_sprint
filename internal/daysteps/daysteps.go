package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	sliseStrings := strings.Split(data, ",")
	if len(sliseStrings) != 2 {
		return 0, 0, fmt.Errorf("ошибочный формат пакета")
	}

	var steps int
	var err error
	steps, err = strconv.Atoi(sliseStrings[0])
	if err != nil {
		return 0, 0, fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("ошибочное количество шагов")
	}

	duration, err := time.ParseDuration(sliseStrings[1])
	if err != nil {
		return 0, 0, fmt.Errorf("conversion error: %w", err)
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil || steps <= 0 {
		return "ошибка"
	}

	distance := float64(steps) * StepLength / 1000

	ccal := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, ccal)
}
