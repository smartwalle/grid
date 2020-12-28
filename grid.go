package zone

type Grid interface {
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

type nGrid struct {
	id   int32
	x    int32
	y    int32
	minX int32
	minY int32
	maxX int32
	maxY int32
}

func NewGrid(id, x, y, minX, minY, maxX, maxY int32) Grid {
	var g = &nGrid{}
	g.id = id
	g.x = x
	g.y = y
	g.minX = minX
	g.minY = minY
	g.maxX = maxX
	g.maxY = maxY
	return g
}

func (this *nGrid) GetId() int32 {
	return this.id
}

func (this *nGrid) GetX() int32 {
	return this.x
}

func (this *nGrid) GetY() int32 {
	return this.y
}

func (this *nGrid) GetMinX() int32 {
	return this.minX
}

func (this *nGrid) GetMaxX() int32 {
	return this.maxX
}

func (this *nGrid) GetMinY() int32 {
	return this.minY
}

func (this *nGrid) GetMaxY() int32 {
	return this.maxY
}

func (this *nGrid) GetWidth() int32 {
	return this.maxX - this.minX + 1
}

func (this *nGrid) GetHeight() int32 {
	return this.maxY - this.minY + 1
}

//func (this *nGrid) String() string {
//	//return fmt.Sprintf("[%3d (%.3d,%.3d)-(%.3d,%.3d)]", this.GetId(), this.GetMinX(), this.GetMinY(), this.GetMaxX(), this.GetMaxY())
//	return fmt.Sprintf("[%3d (%.3d,%.3d)]", this.GetId(), this.GetMinX(), this.GetMinY())
//}
