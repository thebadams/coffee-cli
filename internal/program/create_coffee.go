package program

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type createCoffeeModel struct {
	inputs []textinput.Model
	focus  int
}

func createCoffeePage() createCoffeeModel {
	m := createCoffeeModel{
		inputs: make([]textinput.Model, 3),
	}
	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
		case 1:
			t.Placeholder = "Roaster"
		case 2:
			t.Placeholder = "Origins"
		}
		m.inputs[i] = t
	}
	return m
}

func (m *createCoffeeModel) Init() tea.Cmd {
	return nil
}

func (m *createCoffeeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			new_coffee_page := newCoffeePage()
			return rootProgram().SwitchPage(&new_coffee_page)
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && m.focus == len(m.inputs)-1 {
				new_coffee_page := newCoffeePage()
				return rootProgram().SwitchPage(&new_coffee_page)

			}
			if s == "up" || s == "shift+tab" {
				m.focus--
			} else {
				m.focus++
			}
			if m.focus > len(m.inputs) {
				m.focus = 0
			} else if m.focus < 0 {
				m.focus = len(m.inputs)
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focus {
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				m.inputs[i].Blur()
			}
			return m, tea.Batch(cmds...)
		}
	}
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *createCoffeeModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

func (m *createCoffeeModel) View() string {
	var b strings.Builder
	b.WriteString("Create Coffee Page\n\n")
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}
	return b.String()
}