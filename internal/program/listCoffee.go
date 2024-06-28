package program

func listCoffeeString(p program) string {
	// header
	s := p.programPages[2] + "\n\n"
	for _, coffee := range p.coffees {
		s += coffee + "\n"
	}
	s += "\n Press m to return to main menu\n"
	return s
}
