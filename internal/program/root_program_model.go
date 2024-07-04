package program

import tea "github.com/charmbracelet/bubbletea"

type rootProgramModel struct {
	model tea.Model
	page  int
}

func rootProgram() rootProgramModel {
	main_menu := mainMenu()
	rootModel := &main_menu
	return rootProgramModel{
		model: rootModel,
	}
}

func (r rootProgramModel) Init() tea.Cmd {
	return r.model.Init()
}

func (r rootProgramModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r.model.Update(msg)
}

func (r rootProgramModel) View() string {
	return r.model.View()
}

func (r rootProgramModel) SwitchPage(model tea.Model) (tea.Model, tea.Cmd) {
	r.model = model
	return r.model, r.model.Init()
}
