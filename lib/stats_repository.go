package lib

type Stats struct {
	ID    string
	Stars int
}

type StatsRepository interface {
	Fetch(URL) (Stats, error)
}

type StatsCollectionRepository interface {
	Fetch([]URL) ([]Stats, error)
}
