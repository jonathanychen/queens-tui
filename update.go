package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct {
	t      time.Time
	resets int
}

func (m QueensModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		key := msg.String()
		if key == "ctrl+c" || key == "q" {
			return m, tea.Quit
		}

		if key == "r" {
			m.Reset()
			return m, nil
		}

		if !m.gameOver {
			m.board, _ = m.board.Update(key)

			if m.board.IsGameOver() {
				m.gameOver = true
			}
		}

	case tickMsg:
		if int(msg.resets) == m.resets && !m.gameOver {
			m.timer += 1
		}
		return m, tick(m.resets)
	}

	return m, nil
}

func tick(resets int) tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(tickMsg{t, resets})
	})
}
