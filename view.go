package main

import "fmt"

func (m QueensModel) View() string {
	s := "Queens\n\n"

	s += fmt.Sprintf("%02d:%02d\n", m.timer/60, m.timer%60)

	for i := 0; i < m.board.size+2; i++ {
		s += "-"
	}
	s += fmt.Sprintln()

	for i := 0; i < m.board.size; i++ {
		s += "|"
		for j := 0; j < m.board.size; j++ {
			if i == m.board.cursor.y && j == m.board.cursor.x {
				s += "\033[36m*\033[0m"
			} else {
				s += fmt.Sprintf("%v", colorText(unicodeTile(m.board.grid[i][j]), m.board.colors[i][j]))
			}
		}
		s += "|"
		s += fmt.Sprintln()
	}

	for i := 0; i < m.board.size+2; i++ {
		s += "-"
	}

	s += "\nPress q to quit.\n"

	if m.gameOver {
		s += "You won! Play again? (press r)"
	}

	return s
}

func colorText(input string, color int) string {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	// var Cyan = "\033[36m"
	// var Gray = "\033[37m"
	// var White = "\033[97m"

	colorMap := map[int]string{}
	colorMap[1] = Red
	colorMap[2] = Green
	colorMap[3] = Yellow
	colorMap[4] = Blue
	colorMap[5] = Magenta

	return colorMap[color] + input + Reset
}

func unicodeTile(tile string) string {
	res := ""
	switch tile {
	case " ":
		res = "\u2588"
	case "x":
		res = "\u2A2F"
	case "Q":
		res = "\u2655"
	}
	return res
}
