package program

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type mainMenuModel struct{}

func mainMenu() mainMenuModel {
	return mainMenuModel{}
}

func (m *mainMenuModel) Init() tea.Cmd {
	return nil
}

func (m *mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "l":
			list_coffee_page := listCoffeePage()
			return rootProgram().SwitchPage(&list_coffee_page)
		case "n":
			new_coffee_page := newCoffeePage()
			return rootProgram().SwitchPage(&new_coffee_page)
		}
	}
	return m, nil
}

func (m *mainMenuModel) View() string {
	var b strings.Builder
	b.WriteString("Main Menu\n\n")
	return b.String()
}
