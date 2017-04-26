package lib

type AwesomeStats struct {
	parser       Parser
	repositories map[string]StatsCollectionRepository
	updater      Updater
}

func NewAwesomeStats(parser Parser, repositories map[string]StatsCollectionRepository, updater Updater) AwesomeStats {
	return AwesomeStats{
		parser:       parser,
		repositories: repositories,
		updater:      updater,
	}
}

func (as AwesomeStats) Update(page string) (string, error) {
	urlsDict := as.parser.Parse(page)

	stats := []Stats{}

	for id, urls := range urlsDict {
		repository := as.repositories[id]

		s, err := repository.Fetch(urls)
		if err != nil {
			return "", err
		}

		stats = append(stats, s...)
	}

	result, err := as.updater.Update(page, stats)
	if err != nil {
		return "", err
	}

	// @TODO think about what to do with the result
	return result.page, nil
}
