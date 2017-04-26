package lib

type CompositeUpdater struct {
	updaters []Updater
}

func NewCompositeUpdater(updaters []Updater) CompositeUpdater {
	return CompositeUpdater{
		updaters: updaters,
	}
}

func (u CompositeUpdater) Update(page string, stats []Stats) (updateResult, error) {
	result := updateResult{
		page: page,
	}

	for _, stat := range stats {
		oneStat := []Stats{stat}

		for _, updater := range u.updaters {
			intermediateResult, err := updater.Update(result.page, oneStat)

			if err != nil {
				return result, err
			}

			if intermediateResult.stats.found > 0 {
				result.page = intermediateResult.page
				result.stats.added += intermediateResult.stats.added
				result.stats.found += intermediateResult.stats.found
				result.stats.updated += intermediateResult.stats.updated

				break
			}
		}
	}

	return result, nil
}
