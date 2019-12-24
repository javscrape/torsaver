package torsaver

// Saver ...
type Saver interface {
	Limit(i int)
	CurrentPage() int
	SetPage(int)
	Find(name string) (e error)
	Save(idx int) (e error)
	SaveAll() (e error)
	Download(idx int) (e error)
	DownloadAll() (e error)
	List() []string
}
