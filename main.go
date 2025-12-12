package main

import (
	"fmt"
	"os"
)

// Print error function
func printError() {
	fmt.Println("Error")
}

func validArgs(args []string) bool {
	if len(args) != 10 {
		return false
	}
	for i := 1; i <= 9; i++ {
		if len(args[i]) != 9 {
			return false
		}
		for _, ch := range args[i] {
			if ch != '.' && (ch < '1' || ch > '9') {
				return false
			}
		}
	}
	return true
}

func buildGrid(args []string) [9][9]int {
	var g [9][9]int
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			ch := args[r+1][c]
			if ch == '.' {
				g[r][c] = 0
			} else {
				g[r][c] = int(ch - '0')
			}
		}
	}
	return g
}

func isSafe(g *[9][9]int, r, c, v int) bool {
	for i := 0; i < 9; i++ {
		if g[r][i] == v || g[i][c] == v {
			return false
		}
	}
	br := (r / 3) * 3
	bc := (c / 3) * 3
	for i := br; i < br+3; i++ {
		for j := bc; j < bc+3; j++ {
			if g[i][j] == v {
				return false
			}
		}
	}
	return true
}

func findEmpty(g *[9][9]int) (int, int, bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] == 0 {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func copyGrid(src *[9][9]int) [9][9]int {
	var dst [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			dst[i][j] = src[i][j]
		}
	}
	return dst
}

func solve(g *[9][9]int, count *int, first *[9][9]int) {
	if *count > 1 {
		return
	}

	r, c, ok := findEmpty(g)
	if !ok {
		*count++
		if *count == 1 {
			*first = copyGrid(g)
		}
		return
	}

	for v := 1; v <= 9; v++ {
		if isSafe(g, r, c, v) {
			g[r][c] = v
			solve(g, count, first)
			g[r][c] = 0
		}
	}
}

func sudoku(s []string) {
	grid := buildGrid(s)

	var count int
	var solution [9][9]int

	solve(&grid, &count, &solution)

	if count != 1 {
		printError()
		return
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			fmt.Print(solution[r][c])
			if c < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	if !validArgs(os.Args) {
		printError()
		return
	}
	sudoku(os.Args)
}
