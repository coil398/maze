package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	N = "N"
	S = "S"
	E = "E"
	W = "W"
)

var (
	DIRECTIONS = map[string]int{
		"N": 1,
		"S": 2,
		"E": 4,
		"W": 8,
	}
	DX = map[string]int{
		"E": 1,
		"W": -1,
		"N": 0,
		"S": 0,
	}
	DY = map[string]int{
		"E": 0,
		"W": 0,
		"N": -1,
		"S": 1,
	}
	OPPOSITE = map[string]string{
		"E": "W",
		"W": "E",
		"N": "S",
		"S": "N",
	}
)

func shuffle(d []string) {
	for i := len(d) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func isBetween(i int, min int, max int) bool {
	if i >= min && i <= max {
		return true
	}
	return false
}

func carve_passages_from(cx int, cy int, grid [][]int) {
	directions := []string{"N", "S", "E", "W"}
	shuffle(directions)

	for _, d := range directions {
		nx, ny := cx+DX[d], cy+DY[d]

		if isBetween(ny, 0, len(grid)-1) && isBetween(nx, 0, len(grid[ny])-1) && grid[ny][nx] == 0 {
			grid[cy][cx] |= DIRECTIONS[d]
			grid[ny][nx] |= DIRECTIONS[OPPOSITE[d]]
			carve_passages_from(nx, ny, grid)
		}
	}
}

func debugDisplay(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
}

func display(width int, height int, grid [][]int) {
	fmt.Print(" " + strings.Repeat("_", width*2-1) + "\n")

	for i := 0; i < height; i++ {
		fmt.Print("|")
		for j := 0; j < width; j++ {
			if grid[i][j]&DIRECTIONS["S"] != 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("_")
			}
			if grid[i][j]&DIRECTIONS["E"] != 0 {
				if (grid[i][j]|grid[i][j+1])&DIRECTIONS["S"] != 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("_")
				}
			} else {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func initGrid(width int, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			grid[i][j] = 0
		}
	}

	return grid
}

func main() {
	width := 12
	height := 20

	unixTime := time.Now().Unix()
	rand.Seed(unixTime)

	grid := initGrid(width, height)

	carve_passages_from(0, 0, grid)

	// debugDisplay(grid)

	display(width, height, grid)

	fmt.Printf("width: %d, height: %d, seed: %d", width, height, unixTime)
}
