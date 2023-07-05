package main

func main() {
	g1 := newGrid(10, 10)
	btree(g1)
	g1.print()

	g2 := newGrid(10, 10)
	sidewinder(g2)
	g2.print()
}
