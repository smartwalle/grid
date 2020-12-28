package zone

import (
	"math"
)

type Option func(zone *Zone)

func WithGridMaker(f func(id, x, y, minX, minY, maxX, maxY int32) Grid) Option {
	return func(grid *Zone) {
		grid.nGridFunc = f
	}
}

// WithGridWidth 指定该空间每一个网格的宽度，即每一个网格在 x 轴方向占几个坐标点
func WithGridWidth(width int32) Option {
	return func(grid *Zone) {
		grid.gridWidth = width
	}
}

// WithGridHeight - 指定该空间每一个网格的高度，即每一个网格在 y 轴方向占几个坐标点
func WithGridHeight(height int32) Option {
	return func(grid *Zone) {
		grid.gridHeight = height
	}
}

type Zone struct {
	nGridFunc  func(id, x, y, minX, minY, maxX, maxY int32) Grid
	width      int32
	height     int32
	gridWidth  int32
	gridHeight int32
	gridXCount int32
	gridYCount int32
	grids      []Grid
}

// NewZone 创建新的空间，空间的初始坐标点为 (0, 0)
// width - 指定该空间的宽度，即 x 轴方向有几个坐标点
// height - 指定该空间的高度，即 y 四方向有几个坐标点
//
// 如：创建一个 6 x 6 的空间，每一个网格的大小为 2 x 2：
// var z = NewZone(6, 6, WithGridWidth(2), (WithGridHeight(2))，总共会产生 9 个网格
// 则第一个网格(左上角)的坐标范围为 (0, 0) - (1, 1), 最后一个网格(右下角)的坐标范围为 (4, 4) - (5, 5)
//
// 如：创建一个 5 x 5 的空间，每一个网格的大小为 1 x 1：
// var z = NewZone(5, 5)，总共会产生 25 个网格
// 则第一个网格(左上角)的坐标范围为 (0, 0) - (0, 0), 最后一个网格(右下角)的坐标范围为 (4, 4) - (4, 4)

func NewZone(width, height int32, opts ...Option) *Zone {
	var a = &Zone{}
	a.width = width
	a.height = height

	for _, opt := range opts {
		opt(a)
	}
	if a.nGridFunc == nil {
		a.nGridFunc = NewGrid
	}
	if a.gridWidth <= 0 {
		a.gridWidth = 1
	}
	if a.gridHeight <= 0 {
		a.gridHeight = 1
	}

	a.gridXCount = a.width / a.gridWidth
	a.gridYCount = a.height / a.gridHeight
	if a.width%a.gridWidth > 0 {
		a.gridXCount += 1
	}
	if a.height%a.gridHeight > 0 {
		a.gridYCount += 1
	}
	a.grids = make([]Grid, a.gridXCount*a.gridYCount)

	for y := int32(0); y < a.gridYCount; y++ {
		for x := int32(0); x < a.gridXCount; x++ {
			var gId = x + y*a.gridXCount
			var minX = x * a.gridWidth
			var maxX = (x+1)*a.gridWidth - 1
			var minY = y * a.gridHeight
			var maxY = (y+1)*a.gridHeight - 1
			if maxX >= a.width {
				maxX = a.width - 1
			}
			if maxY >= a.height {
				maxY = a.height - 1
			}
			var grid = a.nGridFunc(gId, x, y, minX, minY, maxX, maxY)
			a.grids[gId] = grid
		}
	}
	return a
}

func (this *Zone) GetWidth() int32 {
	return this.width
}

func (this *Zone) GetHeight() int32 {
	return this.height
}

func (this *Zone) GetGridWidth() int32 {
	return this.gridWidth
}

func (this *Zone) GetGridHeight() int32 {
	return this.gridHeight
}

// GetGridSize 获取 Grid 横/纵的数量
func (this *Zone) GetGridSize() (int32, int32) {
	return this.gridXCount, this.gridYCount
}

// GetGridByPosition 根据坐标点获取 Grid
func (this *Zone) GetGridByPosition(x, y int32) Grid {
	if x < 0 || y < 0 {
		return nil
	}
	if x > this.width || y > this.height {
		return nil
	}

	// 算出坐标所在 Grid
	var gridX = x / this.gridWidth
	var gridY = y / this.gridHeight

	return this.GetGrid(gridX, gridY)
}

// GetGridById 根据 Grid id 获取 Grid
func (this *Zone) GetGridById(gridId int32) Grid {
	if gridId > int32(len(this.grids))-1 || gridId < 0 {
		return nil
	}
	return this.grids[gridId]
}

// GetGrid 根据 Grid 的坐标获取 Grid
func (this *Zone) GetGrid(gridX, gridY int32) Grid {
	var gridId = gridX + gridY*this.gridXCount
	return this.GetGridById(gridId)
}

// GetSurroundGridsByPosition 根据坐标点获取其周边的 Grid 列表
func (this *Zone) GetSurroundGridsByPosition(x, y, round int32) []Grid {
	if x < 0 || y < 0 {
		return nil
	}
	if x > this.width || y > this.height {
		return nil
	}

	// 算出坐标所在 Grid
	var gridX = x / this.gridWidth
	var gridY = y / this.gridHeight

	return this.GetSurroundGrids(gridX, gridY, round)
}

// GetSurroundGridsById 根据 Grid id 获取其周边的 Grid 列表，包含指定 Grid 自身
// round - 指定获取几圈层的网络，为 1 的时候，取其自身加上周边 8 个格子，为 2 的时候，取其自身加上周边 24 个格子，以此类推
func (this *Zone) GetSurroundGridsById(gridId int32, round int32) []Grid {
	if round <= 0 {
		return nil
	}
	var grid = this.GetGridById(gridId)
	if grid == nil {
		return nil
	}

	var gridCount = int(math.Pow(float64(1+round*2), 2))
	var grids = make([]Grid, 0, gridCount)

	var startX = int32(math.Max(float64(grid.GetX()-round), 0))
	var startY = int32(math.Max(float64(grid.GetY()-round), 0))
	var endX = int32(math.Min(float64(grid.GetX()+round), float64(this.gridXCount-1)))
	var endY = int32(math.Min(float64(grid.GetY()+round), float64(this.gridYCount-1)))

	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			var cId = x + y*this.gridXCount
			grids = append(grids, this.grids[cId])
		}
	}
	return grids
}

// GetSurroundGrids 根据 Grid 的坐标获取其周边的 Grid 列表
func (this *Zone) GetSurroundGrids(gridX, gridY, round int32) []Grid {
	var gridId = gridX + gridY*this.gridXCount
	return this.GetSurroundGridsById(gridId, round)
}
