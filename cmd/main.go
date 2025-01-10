package main

import (
	"password-generator/internal/generator"
	"password-generator/internal/storage"
	"strconv"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomTheme struct {
	fyne.Theme
}

func (m *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	default:
		return m.Theme.Color(name, variant)
	}
}

func main() {

	a := app.New()
	a.Settings().SetTheme(&CustomTheme{Theme: theme.DefaultTheme()})
	w := a.NewWindow("Генератор паролей")

	lengthEntry := widget.NewEntry()
	lengthEntry.SetPlaceHolder("Длина пароля (по умолчанию 12)")

	includeSymbols := widget.NewCheck("Символы (!@#$)", nil)
	includeNumbers := widget.NewCheck("Цифры (0-9)", nil)
	includeLower := widget.NewCheck("Строчные буквы (a-z)", nil)
	includeUpper := widget.NewCheck("Заглавные буквы (A-Z)", nil)

	passwordEntry := widget.NewEntry()
	passwordEntry.Disable()
	passwordEntry.SetPlaceHolder("Ваш пароль появится здесь")
	passwordEntry.TextStyle = fyne.TextStyle{Bold: true}

	copyButton := widget.NewButton("Копировать", func() {
		if len(passwordEntry.Text) > 0 {
			w.Clipboard().SetContent(passwordEntry.Text)
		}
	})

	generateButton := widget.NewButton("Сгенерировать", func() {
		length := 12
		if len(lengthEntry.Text) > 0 {
			parsedLength, err := strconv.Atoi(lengthEntry.Text)
			if err != nil || parsedLength <= 0 {
				passwordEntry.SetText("Ошибка: длина должна быть положительным числом")
				return
			}
			if parsedLength > 64 {
				passwordEntry.SetText("Ошибка: длина не может превышать 64 символа")
				return
			}
			length = parsedLength
		}

		config := generator.Config{
			Length:         length,
			IncludeSymbols: includeSymbols.Checked,
			IncludeNumbers: includeNumbers.Checked,
			IncludeLower:   includeLower.Checked,
			IncludeUpper:   includeUpper.Checked,
		}

		password, err := generator.GeneratePassword(config)
		if err != nil {
			passwordEntry.SetText("Ошибка: " + err.Error())
			return
		}

		passwordEntry.SetText(password)

		err = storage.SavePassword("passwords.txt", password)
		if err != nil {
			passwordEntry.SetText("Ошибка сохранения: " + err.Error())
		}
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Настройки генератора паролей"),
		lengthEntry,
		includeSymbols,
		includeNumbers,
		includeLower,
		includeUpper,
		generateButton,
		passwordEntry,
		copyButton,
	))

	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
