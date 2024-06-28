package program

func newCoffeeString(p program) string {
	// Header
	s := p.programPages[1] + "\n\n"

	// Menu
	s += "Press m to return to the main menu"
	return s
}
