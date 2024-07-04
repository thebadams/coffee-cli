package program

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type coffee struct {
	name    string
	roaster string
	origins []string
}
type listCoffeeModel struct {
	coffees []coffee
}

func listCoffeePage() listCoffeeModel {
	coffees := []coffee{
		{"Danche", "Sey Coffee", []string{"Gedeb, Gedeo"}},
		{"Sierra Morena", "Sey Coffee", []string{"Las Mercedes, Palestina, Huila"}},
		{"Finca La Estrella - Late Harvest", "Sey Coffee", []string{"Urrao, Antioquia"}},
		{"Bukeye", "Sey Coffee", []string{"Kayanza"}},
		{"El Cedro", "Sey Coffee", []string{"Purutal, San Agustin, Huila"}},
	}

	return listCoffeeModel{coffees}
}

func (m *listCoffeeModel) Init() tea.Cmd {
	return nil
}

func (m *listCoffeeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "m":
			main_menu := mainMenu()
			return rootProgram().SwitchPage(&main_menu)
		case "n":
			new_coffee_page := newCoffeePage()
			return rootProgram().SwitchPage(&new_coffee_page)

		}
	}
	return m, nil
}

func (m *listCoffeeModel) View() string {
	var b strings.Builder
	b.WriteString("List Coffee Menu\n\n")
	for _, coffee := range m.coffees {
		b.WriteString(coffee.name)
		b.WriteRune('\n')
	}
	return b.String()
}
