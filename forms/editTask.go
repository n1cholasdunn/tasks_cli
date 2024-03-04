package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/n1cholasdunn/tasks_cli/helpers"
)

func EditTaskForm() (map[string]string, error) {
	var (
		editOptions []string
		dueDate     string
		title       string
		notes       string
	)

	editedValues := make(map[string]string)

	isOptionSelected := func(option string) bool {
		for _, selectedOption := range editOptions {
			if selectedOption == option {
				return true
			}
		}
		return false
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("What part of the task/s do you want to modify?").
				Options(
					huh.NewOption("Title", "title"),
					huh.NewOption("Notes", "notes"),
					huh.NewOption("Due Date", "due"),
				).
				Value(&editOptions),
		),
		huh.NewGroup(
			huh.NewInput().Title("Enter new Title").Prompt(":").Value(&title),
		).WithHideFunc(func() bool { return !isOptionSelected("title") }),
		huh.NewGroup(
			huh.NewText().Title("Enter new Notes").Value(&notes),
		).WithHideFunc(func() bool { return !isOptionSelected("notes") }),
		huh.NewGroup(
			huh.NewInput().Title("Enter new Due Date in format YYYY-MM-DD").Prompt(":").Value(&dueDate),
		).WithHideFunc(func() bool { return !isOptionSelected("due") }),
	).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()
	if err != nil {
		return nil, err
	}

	if isOptionSelected("title") {
		editedValues["title"] = title
	}
	if isOptionSelected("notes") {
		editedValues["notes"] = notes
	}
	if isOptionSelected("due") {
		date, err := helpers.ConvertToRFC3339(dueDate)
		if err != nil {
			return nil, err
		}
		editedValues["due"] = date
	}

	return editedValues, nil
}
