package main

import (
	"fmt"
	"github.com/smartwalle/grid"
)

func main() {
	var grid = grid.NewGrid(10, 10, 1, 1)
	var xCount, yCount = grid.GetCellSize()

	for y := int32(0); y < yCount; y++ {
		for x := int32(0); x < xCount; x++ {
			var cell = grid.GetCell(x, y)
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}

	//grid.GetSurroundCellsById(21, 1)
	var cells = grid.GetSurroundCellsByPosition(0, 0, 2)
	fmt.Println(cells)
}
