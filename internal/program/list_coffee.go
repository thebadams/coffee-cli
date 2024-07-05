package program

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thebadams/coffee-cli/db"
)

type listCoffeeModel struct {
	coffees []db.Coffee
	errors  string
}

func connectToDB() tea.Msg {
	database, err := db.Connect()
	if err != nil {
		return err
	}
	data, err := database.GetAllCoffees()
	if err != nil {
		return err
	}
	return data
}

func listCoffeePage() listCoffeeModel {
	return listCoffeeModel{}
}

func (m *listCoffeeModel) Init() tea.Cmd {
	return connectToDB
}

func (m *listCoffeeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		m.coffees = []db.Coffee{{Name: "Coffee", Roaster: "Roaster", Origins: "Origin"}}
		m.errors = msg.Error()
	case []db.Coffee:
		m.coffees = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
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
	b.WriteString("list coffee time\n\n")
	b.WriteString("[m] to return to main menu")
	b.WriteRune('\n')
	b.WriteString("[n] to add new coffee")
	b.WriteRune('\n')
	b.WriteString("[q, esc, or ctrl+c] to quit program")
	b.WriteRune('\n')
	b.WriteRune('\n')
	for _, coffee := range m.coffees {
		b.WriteString(coffee.Name)
		b.WriteRune('\n')
	}
	if m.errors != "" {
		b.WriteString(m.errors)
	}
	return b.String()
}
