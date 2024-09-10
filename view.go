package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// TODO: Eventually needs rework for visuals, table works as a simple solution
// but looks weird. Probably needs some custom table/grid implementation

func (m QueensModel) View() string {
	s := "Queens\n\n"

	s += fmt.Sprintf("Time elapsed: %02d:%02d\n", m.timer/60, m.timer%60)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		BorderStyle(lipgloss.NewStyle()).
		StyleFunc(m.LipglossStyleFunc()).
		Rows(m.LipglossBoardContents(false)...)

	s += t.Render()

	s += "\nPress q to quit.\n"

	if m.gameOver {
		s += "You won! Play again? (press r)"
	}

	return s
}

func (m QueensModel) LipglossStyleFunc() func(row, col int) lipgloss.Style {
	return func(row, col int) lipgloss.Style {
		style := lipgloss.NewStyle().Padding(0, 1)

		if row >= 1 && row <= m.board.size && col >= 0 && col < m.board.size {
			switch m.board.colors[row-1][col] {
			case 1:
				style = style.Background(lipgloss.Color("#FF0000"))
			case 2:
				style = style.Background(lipgloss.Color("#00FF00"))
			case 3:
				style = style.Background(lipgloss.Color("#FFFF00"))
			case 4:
				style = style.Background(lipgloss.Color("#0000FF"))
			case 5:
				style = style.Background(lipgloss.Color("#FF00FF"))
			case 6:
				style = style.Background(lipgloss.Color("#00FFFF"))
			}

			// TODO: Figure out special styling for selected cursor box
			if row-1 == m.board.cursor.y && col == m.board.cursor.x {
				// style = style.Background(lipgloss.Color("#FFFFFF"))
			}
		}

		return style
	}
}

func (m QueensModel) LipglossBoardContents(useUnicode bool) [][]string {
	res := [][]string{}

	for i := 0; i < m.board.size; i++ {
		row := []string{}
		for j := 0; j < m.board.size; j++ {
			if m.board.cursor.x == j && m.board.cursor.y == i && !m.gameOver {
				row = append(row, "*")
			} else {
				if useUnicode {
					row = append(row, unicodeChar(m.board.grid[i][j]))
				} else {
					row = append(row, m.board.grid[i][j])
				}
			}
		}
		res = append(res, row)
	}

	return res
}

func unicodeChar(char string) string {
	switch char {
	case "x":
		return "\u2A2F"
	case "Q":
		return "\u2655"
	default:
		return " "
	}
}
