package program

import "strings"

func newCoffeeString(p program) string {
	// Header
	var b strings.Builder
	b.WriteString(p.programPages[1])
	b.WriteString("\n\n")
	for i := range p.inputs {
		b.WriteString(p.inputs[i].View())
		if i < len(p.inputs)-1 {
			b.WriteRune('\n')
		}
	}
	// Menu
	b.WriteString("Press m to return to main menu")
	return b.String()
}

type coffee struct {
	name    string
	roaster string
	origins []string
}
