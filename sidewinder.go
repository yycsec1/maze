package main

import "math/rand"

func sidewinder(g *grid) {
	for y := 0; y < g.rows; y++ {
		var run []*cell
		for x := 0; x < g.columns; x++ {
			c := g.cells[y][x]
			run = append(run, c)
			atEasternBoundry := c.east == nil
			atNorthernBoundry := c.north == nil

			shouldCloseOut := atEasternBoundry || (!atNorthernBoundry && rand.Intn(2) == 0)

			if shouldCloseOut {
				m := run[rand.Intn(len(run))]
				if m.north != nil {
					m.link(m.north)
				}
				run = []*cell{}
			} else {
				c.link(c.east)
			}
		}
	}
}
