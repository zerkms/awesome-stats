package lib

import (
	"fmt"
	"regexp"
)

type MarkerUpdater struct {
	OpenPrefix  string
	OpenSuffix  string
	ClosePrefix string
	CloseSuffix string
	Formatter   formatter
}

func (m MarkerUpdater) Update(page string, stats []Stats) (updateResult, error) {
	result := updateResult{}

	for _, stat := range stats {
		open := fmt.Sprintf("%s%s%s", m.OpenPrefix, stat.ID, m.OpenSuffix)
		close := fmt.Sprintf("%s%s%s", m.ClosePrefix, stat.ID, m.CloseSuffix)
		statString, err := m.Formatter.format(stat)
		if err != nil {
			return result, err
		}

		reg := fmt.Sprintf("%s.*?%s", regexp.QuoteMeta(open), regexp.QuoteMeta(close))

		r, err := regexp.Compile(reg)
		if err != nil {
			return result, err
		}

		if r.MatchString(page) {
			result.stats.found++
		}

		replacement := fmt.Sprintf("%s%s%s", open, statString, close)

		newPage := r.ReplaceAllString(page, replacement)

		if newPage != page {
			result.stats.updated++
		}

		page = newPage
	}

	result.page = page

	return result, nil
}
