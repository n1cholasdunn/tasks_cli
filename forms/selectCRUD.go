package forms

import (
	"github.com/charmbracelet/huh"
)

func SelectCRUD() (string, error) {
	var selectedOperation string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What task do you want to modify?").
				Options(
					huh.NewOption("Add", "add"),
					huh.NewOption("Delete", "delete"),
					huh.NewOption("Update", "update"),
					huh.NewOption("List", "list"),
				).
				Value(&selectedOperation),
		)).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()
	if err != nil {
		return "", err
	}
	return selectedOperation, nil
}
