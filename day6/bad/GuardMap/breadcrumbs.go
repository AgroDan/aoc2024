package guardmap

/*
 * I initially misread the challenge request, it was looking for _unique_ locations
 * within the traversal. Not necessarily steps. Whoops. So I'm creating this
 */

// need a better variable to store
//var exists = struct{}{}

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
	_, e := b.crumb[c]
	return e
}

func (b *Breadcrumb) GetDir(c Coord) int {
	return b.crumb[c]
}

func (b Breadcrumb) Amount() int {
	return len(b.crumb)
}
