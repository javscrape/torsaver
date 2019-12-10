package torsaver

type Saver interface {
	Limit(i int64)
	Find(name string) (e error)
	Save(idx int, path string) (e error)
	SaveAll(path string) (e error)
	List() []string
}
