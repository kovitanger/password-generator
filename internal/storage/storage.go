package storage

import (
	"fmt"
	"os"
	"time"
)

func SavePassword(filename string, password string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	entry := fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), password)
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	return nil
}
