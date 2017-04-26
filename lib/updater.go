package lib

type updateStats struct {
	found   int
	added   int
	updated int
}

type updateResult struct {
	page  string
	stats updateStats
}

type Updater interface {
	Update(page string, stats []Stats) (updateResult, error)
}
