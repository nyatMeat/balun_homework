package logger

import (
	"fmt"
)

const (
	stringLevelDebug   = "debug"
	stringLevelInfo    = "info"
	stringLevelWarning = "warning"
	stringLevelError   = "error"
)

var logLevelMap = map[string]Level{
	stringLevelDebug:   LevelDebug,
	stringLevelInfo:    LevelInfo,
	stringLevelWarning: LevelWarn,
	stringLevelError:   LevelError,
}

func ConvertHumanReadableLevelToLevel(level string) (Level, error) {
	v, ok := logLevelMap[level]

	if !ok {
		return LevelDebug, fmt.Errorf("incorrect logging level")
	}

	return v, nil
}
