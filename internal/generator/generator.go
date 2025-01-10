package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Config struct {
	Length         int
	IncludeSymbols bool
	IncludeNumbers bool
	IncludeLower   bool
	IncludeUpper   bool
}

func GeneratePassword(config Config) (string, error) {
	if config.Length <= 0 {
		return "", fmt.Errorf("длина пароля должна быть больше 0")
	}

	rand.Seed(time.Now().UnixNano())

	var charset strings.Builder

	if config.IncludeLower {
		charset.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	if config.IncludeUpper {
		charset.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if config.IncludeNumbers {
		charset.WriteString("0123456789")
	}
	if config.IncludeSymbols {
		charset.WriteString("!@#$%^&*()-_=+[]{}|;:,.<>?/`~")
	}

	if charset.Len() == 0 {
		return "", fmt.Errorf("не указан ни один допустимый набор символов")
	}

	var password strings.Builder
	for i := 0; i < config.Length; i++ {
		password.WriteByte(charset.String()[rand.Intn(charset.Len())])
	}

	return password.String(), nil
}
