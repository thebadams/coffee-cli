package program

func mainMenuString(p program) string {
	// HEADER
	s := p.programPages[0] + "\n\n"

	// Menu
	s += "(Press n to create a new coffee"

	return s
}

func mainMenu() page {
	return page{
		name:  "Main Menu",
		index: 0,
	}
}
