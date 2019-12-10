package torsaver

type Saver interface {
	Limit(i int64)
	Find(name string) (e error)
}
