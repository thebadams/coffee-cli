package program

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type program struct {
	coffees      []coffee
	inputs       []textinput.Model
	programPages []string
	programState int
	inputIndex   int
}

type page struct {
	name  string
	index int
}

func defaultProgram() program {
	p := program{
		programPages: []string{"main menu", "new coffee", "list coffee"},
		coffees: []coffee{
			{"Munyinya Hill, Lot #2", "Sey Coffee", []string{}},
			{"Finca La Estrella - Late Harvest", "Sey Coffee", []string{}},
			{"El Cedro", "Sey Coffee", []string{}},
			{"Sierra Morena", "Sey Coffee", []string{}},
		},
		inputs: make([]textinput.Model, 3),
	}

	var t textinput.Model
	for i := range p.inputs {
		t = textinput.New()
		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
		case 1:
			t.Placeholder = "Roaster"
		case 2:
			t.Placeholder = "Origin(s)"
		}
		p.inputs[i] = t
	}
	return p
}

func (p *program) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(p.inputs))
	for i := range p.inputs {
		p.inputs[i], cmds[i] = p.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

func (p program) Init() tea.Cmd {
	return nil
}

func (p program) View() string {
	var s string
	switch p.programState {
	case 0:
		s = mainMenuString(p)
	case 1:
		s = newCoffeeString(p)
	case 2:
		s = listCoffeeString(p)
	}

	s += footerString()
	return s
}

func (p program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// Universal Key Inputs
		case "ctrl+c", "esc":
			return p, tea.Quit
		case "ctrl+m":
			if p.programState != 0 {
				p.programState = 0
			}
		case "ctrl+n":
			if p.programState != 1 {
				p.programState = 1
			}
		case "ctrl+l":
			if p.programState != 2 {
				p.programState = 2
			}
		default:
			if p.programState == 2 {
				cmd := p.updateInputs(msg)
				return p, cmd

			}
		}
	}
	return p, nil
}

func CreateProgram() *tea.Program {
	p := tea.NewProgram(defaultProgram())

	return p
}
