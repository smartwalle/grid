package grid

type Cell struct {
	id   int32
	x    int32
	y    int32
	minX int32
	minY int32
	maxX int32
	maxY int32
}

func NewCell(id, x, y, minX, minY, maxX, maxY int32) *Cell {
	var cell = &Cell{}
	cell.id = id
	cell.x = x
	cell.y = y
	cell.minX = minX
	cell.minY = minY
	cell.maxX = maxX
	cell.maxY = maxY
	return cell
}

func (this *Cell) GetId() int32 {
	return this.id
}

func (this *Cell) GetX() int32 {
	return this.x
}

func (this *Cell) GetY() int32 {
	return this.y
}

func (this *Cell) GetMinX() int32 {
	return this.minX
}

func (this *Cell) GetMaxX() int32 {
	return this.maxX
}

func (this *Cell) GetMinY() int32 {
	return this.minY
}

func (this *Cell) GetMaxY() int32 {
	return this.maxY
}

func (this *Cell) GetWidth() int32 {
	return this.maxX - this.minX + 1
}

func (this *Cell) GetHeight() int32 {
	return this.maxY - this.minY + 1
}

//func (this *Cell) String() string {
//	//return fmt.Sprintf("[%3d (%.3d,%.3d)-(%.3d,%.3d)]", this.GetId(), this.GetMinX(), this.GetMinY(), this.GetMaxX(), this.GetMaxY())
//	return fmt.Sprintf("[%3d (%.3d,%.3d)]", this.GetId(), this.GetMinX(), this.GetMinY())
//}
