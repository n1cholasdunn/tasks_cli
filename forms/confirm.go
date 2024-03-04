package forms

import "github.com/charmbracelet/huh"

func ConfirmDeleteForm() (bool, error) {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Are you sure you want to delete?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&confirm),
		)).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()
	if err != nil {
		return false, err
	}
	return confirm, nil
}

func ConfirmEditForm() (bool, error) {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Are you okay with these changes?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&confirm),
		)).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()
	if err != nil {
		return false, err
	}
	return confirm, nil
}
