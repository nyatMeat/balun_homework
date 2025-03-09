package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var sizeUnits = map[string]int{
	"B":  1,
	"KB": 1 << 10, // 1024
	"MB": 1 << 20, // 1048576
	"GB": 1 << 30, // 1073741824
	"TB": 1 << 40, // 1099511627776
	"PB": 1 << 50, // 1125899906842624
	"EB": 1 << 60, // 1152921504606846976
}

func ParseHumanReadableSize(sizeStr string) (int, error) {
	// Убираем пробелы и приводим к верхнему регистру
	cleanStr := strings.ToUpper(strings.TrimSpace(sizeStr))

	// Регулярка для разделения числа и единицы измерения
	re := regexp.MustCompile(`^\s*(\d+)\s*([A-Za-z]+)\s*$`)
	matches := re.FindStringSubmatch(cleanStr)

	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid format")
	}

	// Парсим числовую часть
	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("invalid number format")
	}

	if num < 0 {
		return 0, fmt.Errorf("negative size value")
	}

	// Нормализуем единицу измерения
	unit := strings.TrimSpace(matches[2])

	if unit == "" {
		return 0, fmt.Errorf("missing size unit")
	}

	// Проверяем поддерживаемые единицы измерения
	multiplier, exists := sizeUnits[unit]

	if !exists {
		// Пробуем найти вариант без окончания B (KB -> K)
		if multiplier, exists = sizeUnits[unit+"B"]; !exists {
			return 0, fmt.Errorf("unsupported unit: %s", unit)
		}
	}

	return num * multiplier, nil
}
