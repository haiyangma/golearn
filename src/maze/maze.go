package main

import (
	"fmt"
	"os"
	"time"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j  int
	value int
}

func sniff(maze *[][]int, p point, steps *[][]int) []point {
	var q []point
	var x, y int
	if p.i > 0 {
		x = p.i - 1
		y = p.j
		if (*maze)[x][y] == 0 {
			if (x != 0 || y != 0) && (*steps)[x][y] == 0 {
				(*steps)[x][y] = (*steps)[p.i][p.j] + 1
				q = append(q, point{i: x, j: y})
			}
		}
	}
	if p.j > 0 {
		x = p.i
		y = p.j - 1
		if (*maze)[x][y] == 0 {
			if (x != 0 || y != 0) && (*steps)[x][y] == 0 {
				(*steps)[x][y] = (*steps)[p.i][p.j] + 1
				q = append(q, point{i: x, j: y})
			}
		}
	}
	if p.i < len(*maze)-1 {
		x = p.i + 1
		y = p.j
		if (*maze)[x][y] == 0 {
			if (x != 0 || y != 0) && (*steps)[x][y] == 0 {
				(*steps)[x][y] = (*steps)[p.i][p.j] + 1
				q = append(q, point{i: x, j: y})
			}
		}
	}
	if p.j < len((*maze)[0])-1 {
		x = p.i
		y = p.j + 1
		if (*maze)[x][y] == 0 {
			if (x != 0 || y != 0) && (*steps)[x][y] == 0 {
				(*steps)[x][y] = (*steps)[p.i][p.j] + 1
				q = append(q, point{i: x, j: y})
			}
		}
	}
	return q
}

func walk(maze *[][]int, start, end point) {
	steps := make([][]int, len(*maze))
	for i := range *maze {
		steps[i] = make([]int, len((*maze)[0]))
	}
	Q := []point{start}
	for {
		if len(Q) == 0 {
			break
		}
		q := sniff(maze, Q[0], &steps)
		for _, p := range q {
			Q = append(Q, p)
		}
		Q = Q[1:]
	}
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

func main() {
	t1 := time.Now()
	maze := readMaze("src/maze/maze.in")

	walk(&maze, point{0, 0, 0}, point{len(maze) - 1, len(maze[0]) - 1, 0})
	fmt.Println("used time %d ", time.Since(t1).Nanoseconds())
}
