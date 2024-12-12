package utils

// This is a series of breadcrumbs to use for a rectangular map.
// It is meant to record all previously-visited positions on the
// map, as well as the direction you were facing when doing so.

type Breadcrumb struct {
	crumb map[Coord]int
}

func NewBreadcrumb() *Breadcrumb {
	b := &Breadcrumb{}
	b.crumb = make(map[Coord]int)
	return b
}

func (b *Breadcrumb) Add(c Coord, dir int) {
	b.crumb[c] = dir
}

func (b *Breadcrumb) Remove(c Coord) {
	delete(b.crumb, c)
}

func (b *Breadcrumb) Contains(c Coord) bool {
	_, exists := b.crumb[c]
	return exists
}

func (b *Breadcrumb) GetDir(c Coord) int {
	return b.crumb[c]
}

func (b *Breadcrumb) Amount() int {
	return len(b.crumb)
}

func (b *Breadcrumb) List() map[Coord]int {
	return b.crumb
}

func (b Breadcrumb) DeepCopy() *Breadcrumb {
	newBC := NewBreadcrumb()
	for k, v := range b.crumb {
		newBC.Add(k, v)
	}
	return newBC
}
