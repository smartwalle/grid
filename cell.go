package grid

type Cell interface {
	GetId() int32

	GetX() int32

	GetY() int32

	GetMinX() int32

	GetMaxX() int32

	GetMinY() int32

	GetMaxY() int32

	GetWidth() int32

	GetHeight() int32
}

type nCell struct {
	id   int32
	x    int32
	y    int32
	minX int32
	minY int32
	maxX int32
	maxY int32
}

func NewCell(id, x, y, minX, minY, maxX, maxY int32) Cell {
	var cell = &nCell{}
	cell.id = id
	cell.x = x
	cell.y = y
	cell.minX = minX
	cell.minY = minY
	cell.maxX = maxX
	cell.maxY = maxY
	return cell
}

func (this *nCell) GetId() int32 {
	return this.id
}

func (this *nCell) GetX() int32 {
	return this.x
}

func (this *nCell) GetY() int32 {
	return this.y
}

func (this *nCell) GetMinX() int32 {
	return this.minX
}

func (this *nCell) GetMaxX() int32 {
	return this.maxX
}

func (this *nCell) GetMinY() int32 {
	return this.minY
}

func (this *nCell) GetMaxY() int32 {
	return this.maxY
}

func (this *nCell) GetWidth() int32 {
	return this.maxX - this.minX + 1
}

func (this *nCell) GetHeight() int32 {
	return this.maxY - this.minY + 1
}

//func (this *nCell) String() string {
//	//return fmt.Sprintf("[%3d (%.3d,%.3d)-(%.3d,%.3d)]", this.GetId(), this.GetMinX(), this.GetMinY(), this.GetMaxX(), this.GetMaxY())
//	return fmt.Sprintf("[%3d (%.3d,%.3d)]", this.GetId(), this.GetMinX(), this.GetMinY())
//}
