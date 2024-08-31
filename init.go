package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() QueensModel {
	model := QueensModel{}
	model.Reset()
	return model
}

func (m QueensModel) Init() tea.Cmd {
	return tick(1)
}
