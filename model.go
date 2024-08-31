package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	numBoards int32 = 1
)

type QueensModel struct {
	board    Board
	timer    int
	gameOver bool
	resets   int
	// windowWidth  int // TODO: modify view based on window size
	// windowHeight int
}

func (m *QueensModel) Reset() {
	startingBoardNum := rand.Int31n(numBoards) + 1
	boardFileName := fmt.Sprintf("boards/board%v.json", startingBoardNum)

	boardFileJSON, err := os.ReadFile(boardFileName)
	if err != nil {
		panic(err)
	}

	var boardFileGrid [][]int
	if err := json.Unmarshal(boardFileJSON, &boardFileGrid); err != nil {
		panic(err)
	}

	grid := [][]string{}
	colors := [][]int{}

	size := len(boardFileGrid)

	for i := 0; i < size; i++ {
		gridRow := []string{}
		colorRow := []int{}

		for j := 0; j < size; j++ {
			gridRow = append(gridRow, " ")
			colorRow = append(colorRow, boardFileGrid[i][j])
		}
		grid = append(grid, gridRow)
		colors = append(colors, colorRow)
	}

	m.board = Board{
		grid:   grid,
		colors: colors,
		size:   size,
		cursor: Pos{
			x: 0,
			y: 0,
		},
	}
	m.timer = 0
	m.gameOver = false
	m.resets += 1
}

type Pos struct {
	x int
	y int
}

type Board struct {
	grid   [][]string
	colors [][]int
	size   int
	cursor Pos
}

func (b Board) Update(key string) (Board, tea.Cmd) {
	switch key {
	case "left", "h":
		if b.cursor.x > 0 {
			b.cursor.x--
		}

	case "right", "l":
		if b.cursor.x < b.size-1 {
			b.cursor.x++
		}

	case "up", "k":
		if b.cursor.y > 0 {
			b.cursor.y--
		}

	case "down", "j":
		if b.cursor.y < b.size-1 {
			b.cursor.y++
		}

	case "enter", " ":
		x := b.cursor.x
		y := b.cursor.y

		switch b.grid[y][x] {
		case " ":
			b.grid[y][x] = "x"
		case "x":
			b.grid[y][x] = "Q"
		case "Q":
			b.grid[y][x] = " "
		}

	case "c":
		for i := 0; i < b.size; i++ {
			for j := 0; j < b.size; j++ {
				b.grid[i][j] = " "
			}
		}
	}
	return b, nil
}

func (b Board) IsGameOver() bool {
	queenPositions := []Pos{}

	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.grid[i][j] == "Q" {
				queenPositions = append(queenPositions, Pos{x: j, y: i})
			}
		}
	}

	// Check if more queens than size of board
	if len(queenPositions) != b.size {
		return false
	}

	// Check if there are diagonally adjacent queens
	diags := []Pos{{x: 1, y: 1}, {x: 1, y: -1}, {x: -1, y: 1}, {x: -1, y: -1}}

	for _, p := range queenPositions {
		for _, d := range diags {
			nx, ny := p.x+d.x, p.y+d.y
			if nx >= 0 && nx < b.size && ny >= 0 && ny < b.size && b.grid[ny][nx] == "Q" {
				return false
			}
		}
	}

	// Check if there is more than one queen per col
	for col := 0; col < b.size; col++ {
		count := 0
		for row := 0; row < b.size; row++ {
			if b.grid[row][col] == "Q" {
				count++
			}
		}
		if count != 1 {
			return false
		}
	}

	// Check if there is more than one queen per row
	for row := 0; row < b.size; row++ {
		count := 0
		for col := 0; col < b.size; col++ {
			if b.grid[row][col] == "Q" {
				count++
			}
		}
		if count != 1 {
			return false
		}
	}

	directions := []Pos{{x: 0, y: 1}, {x: 0, y: -1}, {x: 1, y: 0}, {x: -1, y: 0}}

	for _, q := range queenPositions {
		startColor := b.colors[q.y][q.x]

		queue := []Pos{q}

		seen := map[Pos]bool{}
		seen[q] = true

		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]

			if c.x != q.x && c.y != q.y && b.grid[c.y][c.x] == "Q" {
				return false
			}

			for _, d := range directions {
				nx, ny := c.x+d.x, c.y+d.y
				if nx >= 0 && nx < b.size && ny >= 0 && ny < b.size && b.colors[ny][nx] == startColor {
					if _, ok := seen[Pos{x: nx, y: ny}]; !ok {
						queue = append(queue, Pos{x: nx, y: ny})
						seen[Pos{x: nx, y: ny}] = true
					}
				}
			}

		}

	}

	return true
}
