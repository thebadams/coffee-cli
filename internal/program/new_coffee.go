package program

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type newCoffeeModel struct{}

func newCoffeePage() newCoffeeModel {
	return newCoffeeModel{}
}

func (m *newCoffeeModel) Init() tea.Cmd {
	return nil
}

func (m *newCoffeeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "m":
			main_menu := mainMenu()
			return rootProgram().SwitchPage(&main_menu)
		case "l":
			list_coffee_page := listCoffeePage()
			return rootProgram().SwitchPage(&list_coffee_page)
		case "c":
			create_coffee_page := createCoffeePage()
			return rootProgram().SwitchPage(&create_coffee_page)
		}
	}
	return m, nil
}

func (m *newCoffeeModel) View() string {
	var b strings.Builder
	b.WriteString("new coffee time!")
	b.WriteRune('\n')
	b.WriteRune('\n')
	b.WriteString("[l] to list coffees")
	b.WriteRune('\n')
	b.WriteString("[m] to return to main menu")
	b.WriteRune('\n')
	b.WriteString("[c] to create a new coffee")
	b.WriteRune('\n')
	b.WriteString("[q, esc, or ctrl+c] to quit program")

	return b.String()
}
