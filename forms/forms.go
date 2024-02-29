package forms

import (
	"log"

	"github.com/charmbracelet/huh"
)

func Form() {
	type Task struct {
		Title string
		Notes *string
		Due   *string
	}

	var addNewTask string
	var addingTask []string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Do you want to add a new task?").
				Options(
					huh.NewOption("Yes", "yes"),
					huh.NewOption("No", "no"),
				).
				Value(&addNewTask),
			huh.NewMultiSelect[string]().
				Title("Are you done adding Tasks? Y/N").
				Options(
					huh.NewOption("Yes", "yes"),
					huh.NewOption("No", "no"),
				).Value(&addingTask),
		),
	).WithTheme(huh.ThemeCatppuccin())
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
}
