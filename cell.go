package main

type cell struct {
	row    int
	column int
	north  *cell
	south  *cell
	east   *cell
	west   *cell
	links  map[*cell]bool
}

func newCell(row, column int) *cell {
	return &cell{row: row, column: column, links: make(map[*cell]bool)}
}

func (c *cell) link(n *cell) {
	c.links[n] = true
	n.links[c] = true
}

func (c *cell) unlink(n *cell, biDir bool) {
	delete(c.links, n)
	delete(n.links, c)
}

func (c *cell) linkedCells() []*cell {
	var r []*cell
	for k := range c.links {
		r = append(r, k)
	}
	return r
}

func (c *cell) isLinked(n *cell) bool {
	return c.links[n]
}

func (c *cell) neighbors() []*cell {
	var n []*cell
	if c.north != nil {
		n = append(n, c.north)
	}
	if c.south != nil {
		n = append(n, c.south)
	}
	if c.east != nil {
		n = append(n, c.east)
	}
	if c.west != nil {
		n = append(n, c.west)
	}
	return n
}
