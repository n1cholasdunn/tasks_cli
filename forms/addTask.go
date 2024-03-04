package forms

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/n1cholasdunn/tasks_cli/helpers"
)

func AddTaskForm() (string, string, string, error) {
	var (
		title      string
		notes      string
		dueDate    string
		customDate string
	)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Enter Task Title").Prompt("-->").Value(&title),
			huh.NewText().Title("Enter Notes for the task (press enter to skip)").Value(&notes),
			huh.NewSelect[string]().
				Title("What do you want as the due date?").
				Options(
					huh.NewOption("Today", helpers.GetTodayAsRFC3339()),
					huh.NewOption("Tomorrow", helpers.GetTomorrowAsRFC3339()),
					huh.NewOption("Custom", "custom"),
				).Value(&dueDate),
		),
		huh.NewGroup(
			huh.NewInput().Title("Enter Due Date in format YYYY-MM-DD").Prompt(":").Validate(helpers.ValidateDate).Value(&customDate),
		).WithHideFunc(func() bool { return dueDate != "custom" }),
	)

	err := form.Run()
	if err != nil {
		return "", "", "", err
	}

	if dueDate == "custom" && customDate != "" {
		convertedDate, err := helpers.ConvertToRFC3339(customDate)
		if err != nil {
			return "", "", "", fmt.Errorf("error converting custom date to RFC3339: %v", err)
		}
		dueDate = convertedDate
	}

	return title, notes, dueDate, nil
}
