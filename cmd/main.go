package main

import (
	"fmt"
	"github.com/smartwalle/zone"
)

func main() {
	var nZone = zone.NewZone(10, 10)
	var xCount, yCount = nZone.GetGridSize()

	for y := int32(0); y < yCount; y++ {
		for x := int32(0); x < xCount; x++ {
			var grid = nZone.GetGrid(x, y)
			fmt.Print(grid, " ")
		}
		fmt.Println()
	}

	var grids = nZone.GetSurroundGridsByPosition(0, 0, 2)
	fmt.Println(grids)
}
