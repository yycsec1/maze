package main

import "math/rand"

func btree(g *grid) {
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.columns; x++ {
			c := g.cells[y][x]
			var n []*cell
			if c.north != nil {
				n = append(n, c.north)
			}
			if c.east != nil {
				n = append(n, c.east)
			}
			if len(n) > 0 {
				c.link(n[rand.Intn(len(n))])
			}
		}
	}
}
