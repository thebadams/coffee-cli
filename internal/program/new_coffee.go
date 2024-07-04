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
		case "ctrl+c", "esc":
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
	b.WriteString("New Coffee Page")
	return b.String()
}
