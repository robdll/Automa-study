package entita

type AutomaGroup struct {
	Distanza int
	Automi   []*Automa
}

func GetManhattanDistance(a, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
