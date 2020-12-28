package grid

import (
	"math"
)

type Option func(grid *Grid)

func WithCellMaker(f func(id, x, y, minX, minY, maxX, maxY int32) Cell) Option {
	return func(grid *Grid) {
		grid.nCellFunc = f
	}
}

func WithCellWidth(width int32) Option {
	return func(grid *Grid) {
		grid.cellWidth = width
	}
}

func WithCellHeight(height int32) Option {
	return func(grid *Grid) {
		grid.cellHeight = height
	}
}

type Grid struct {
	nCellFunc  func(id, x, y, minX, minY, maxX, maxY int32) Cell
	width      int32
	height     int32
	cellWidth  int32
	cellHeight int32
	cellXCount int32
	cellYCount int32
	cells      []Cell
}

// NewGrid 创建新的网格，网络的初始坐标点为 (0, 0)
// width - 指定该网格的宽度，即 x 轴方向有几个坐标点
// height - 指定该网格的高度，即 y 四方向有几个坐标点
// cellWidth - 指定该网格每一个格子的宽度，即每一个格子在 x 轴方向占几个坐标点
// cellHeight - 指定该网络每一个格子的高度，即每一个格子在 y 轴方向点几个坐标点
//
// 如：创建一个 6 x 6 的网格，每一个格子的大小为 2 x 2：
// var g = NewGrid(6, 6, 2, 2)，总共会产生 9 个格子
// 则第一个格子(左上角)的坐标范围为 (0, 0) - (1, 1), 最后一个格子(右下角)的坐标范围为 (4, 4) - (5, 5)
//
// 如：创建一个 5 x 5 的网络，每一个格子的大小为 1 x 1：
// var g = NewGrid(5, 5, 1, 1)，总共会产生 25 个格子
// 则第一个格子(左上角)的坐标范围为 (0, 0) - (0, 0), 最后一个格子(右下角)的坐标范围为 (4, 4) - (4, 4)

func NewGrid(width, height int32, opts ...Option) *Grid {
	var a = &Grid{}
	a.width = width
	a.height = height

	for _, opt := range opts {
		opt(a)
	}
	if a.nCellFunc == nil {
		a.nCellFunc = NewCell
	}
	if a.cellWidth <= 0 {
		a.cellWidth = 1
	}
	if a.cellHeight <= 0 {
		a.cellHeight = 1
	}

	a.cellXCount = a.width / a.cellWidth
	a.cellYCount = a.height / a.cellHeight
	if a.width%a.cellWidth > 0 {
		a.cellXCount += 1
	}
	if a.height%a.cellHeight > 0 {
		a.cellYCount += 1
	}
	a.cells = make([]Cell, a.cellXCount*a.cellYCount)

	for y := int32(0); y < a.cellYCount; y++ {
		for x := int32(0); x < a.cellXCount; x++ {
			var cId = x + y*a.cellXCount
			var minX = x * a.cellWidth
			var maxX = (x+1)*a.cellWidth - 1
			var minY = y * a.cellHeight
			var maxY = (y+1)*a.cellHeight - 1
			if maxX >= a.width {
				maxX = a.width - 1
			}
			if maxY >= a.height {
				maxY = a.height - 1
			}
			var cell = a.nCellFunc(cId, x, y, minX, minY, maxX, maxY)
			a.cells[cId] = cell
		}
	}
	return a
}

func (this *Grid) GetWidth() int32 {
	return this.width
}

func (this *Grid) GetHeight() int32 {
	return this.height
}

func (this *Grid) GetCellWidth() int32 {
	return this.cellWidth
}

func (this *Grid) GetCellHeight() int32 {
	return this.cellHeight
}

// GetCellSize 获取 nCell 横/纵的数量
func (this *Grid) GetCellSize() (int32, int32) {
	return this.cellXCount, this.cellYCount
}

// GetCellByPosition 根据坐标点获取 nCell
func (this *Grid) GetCellByPosition(x, y int32) Cell {
	if x < 0 || y < 0 {
		return nil
	}
	if x > this.width || y > this.height {
		return nil
	}

	// 算出坐标所在 nCell
	var CellX = x / this.cellWidth
	var cellY = y / this.cellHeight

	return this.GetCell(CellX, cellY)
}

// GetCellById 根据 nCell id 获取 nCell
func (this *Grid) GetCellById(cellId int32) Cell {
	if cellId > int32(len(this.cells))-1 || cellId < 0 {
		return nil
	}
	return this.cells[cellId]
}

// GetCell 根据 nCell 的坐标获取 nCell
func (this *Grid) GetCell(cellX, CellY int32) Cell {
	var cellId = cellX + CellY*this.cellXCount
	return this.GetCellById(cellId)
}

// GetSurroundCellsByPosition 根据坐标点获取其周边的 nCell 列表
func (this *Grid) GetSurroundCellsByPosition(x, y, round int32) []Cell {
	if x < 0 || y < 0 {
		return nil
	}
	if x > this.width || y > this.height {
		return nil
	}

	// 算出坐标所在 nCell
	var CellX = x / this.cellWidth
	var cellY = y / this.cellHeight

	return this.GetSurroundCells(CellX, cellY, round)
}

// GetSurroundCellsById 根据 nCell id 获取其周边的 nCell 列表，包含指定 nCell 自身
// round - 指定获取几圈层的格子，为 1 的时候，取其自身加上周边 8 个格子，为 2 的时候，取其自身加上周边 24 个格子，以此类推
func (this *Grid) GetSurroundCellsById(cellId int32, round int32) []Cell {
	if round <= 0 {
		return nil
	}
	var cell = this.GetCellById(cellId)
	if cell == nil {
		return nil
	}

	var cellCount = int(math.Pow(float64(1+round*2), 2))
	var cells = make([]Cell, 0, cellCount)

	var startX = int32(math.Max(float64(cell.GetX()-round), 0))
	var startY = int32(math.Max(float64(cell.GetY()-round), 0))
	var endX = int32(math.Min(float64(cell.GetX()+round), float64(this.cellXCount-1)))
	var endY = int32(math.Min(float64(cell.GetY()+round), float64(this.cellYCount-1)))

	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			var cId = x + y*this.cellXCount
			cells = append(cells, this.cells[cId])
		}
	}
	return cells
}

// GetSurroundCells 根据 nCell 的坐标获取其周边的 nCell 列表
func (this *Grid) GetSurroundCells(cellX, CellY, round int32) []Cell {
	var cellId = cellX + CellY*this.cellXCount
	return this.GetSurroundCellsById(cellId, round)
}
