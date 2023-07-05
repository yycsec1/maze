package main

import "fmt"

type grid struct {
	rows    int
	columns int
	cells   [][]*cell
}

func newGrid(rows, columns int) *grid {
	g := &grid{rows: rows, columns: columns}
	g.cells = make([][]*cell, rows)
	for y := 0; y < g.rows; y++ {
		g.cells[y] = make([]*cell, columns)
		for x := 0; x < g.columns; x++ {
			g.cells[y][x] = newCell(y, x)
		}
	}
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.columns; x++ {
			c := g.cells[y][x]
			if y > 0 {
				c.north = g.cells[y-1][x]
			}
			if y < g.rows-1 {
				c.south = g.cells[y+1][x]
			}
			if x > 0 {
				c.west = g.cells[y][x-1]
			}
			if x < g.columns-1 {
				c.east = g.cells[y][x+1]
			}
		}
	}
	return g
}

func (g *grid) toString() string {
	output := "+"
	for x := 0; x < g.columns; x++ {
		output += "---+"
	}
	output += "\n"
	for y := 0; y < g.rows; y++ {
		top := "|"
		bottom := "+"
		for x := 0; x < g.columns; x++ {
			c := g.cells[y][x]
			body := "   "
			east := "|"
			if c.isLinked(c.east) {
				east = " "
			}
			top += body + east
			south := "---"
			if c.isLinked(c.south) {
				south = "   "
			}
			bottom += south + "+"
		}
		output += top + "\n" + bottom + "\n"
	}
	return output
}

func (g *grid) print() {
	fmt.Println(g.toString())
}
