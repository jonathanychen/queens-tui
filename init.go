package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() QueensModel {
	model := QueensModel{}
	model.Reset()
	return model

	// startingBoardNum := rand.Int31n(numBoards) + 1
	// boardFileName := fmt.Sprintf("boards/board%v.json", startingBoardNum)

	// boardFileJSON, err := os.ReadFile(boardFileName)
	// if err != nil {
	// 	panic(err)
	// }

	// var boardFileGrid [][]int
	// if err := json.Unmarshal(boardFileJSON, &boardFileGrid); err != nil {
	// 	panic(err)
	// }

	// grid := [][]string{}
	// colors := [][]int{}

	// size := len(boardFileGrid)

	// for i := 0; i < size; i++ {
	// 	gridRow := []string{}
	// 	colorRow := []int{}

	// 	for j := 0; j < size; j++ {
	// 		gridRow = append(gridRow, " ")
	// 		colorRow = append(colorRow, boardFileGrid[i][j])
	// 	}
	// 	grid = append(grid, gridRow)
	// 	colors = append(colors, colorRow)
	// }

	// board := Board{
	// 	grid:   grid,
	// 	colors: colors,
	// 	size:   size,
	// 	cursor: Pos{
	// 		x: 0,
	// 		y: 0,
	// 	},
	// }

	// return QueensModel{
	// 	board:    board,
	// 	timer:    0,
	// 	gameOver: false,
	// }
}

func (m QueensModel) Init() tea.Cmd {
	return tick(1)
}
