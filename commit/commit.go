package commit

type Line struct {
	Location int
	val      string
}

func (l Line) String() string {
	return l.val
}

type Change struct {
	Added   []Line
	Removed []Line
}
