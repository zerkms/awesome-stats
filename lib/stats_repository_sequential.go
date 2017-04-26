package lib

type SequentialRepository struct {
	Repository StatsRepository
}

func (r SequentialRepository) Fetch(urls []URL) ([]Stats, error) {
	result := make([]Stats, 0, len(urls))

	for _, url := range urls {
		stat, err := r.Repository.Fetch(url)
		if err != nil {
			return result, err
		}

		result = append(result, stat)
	}

	return result, nil
}
