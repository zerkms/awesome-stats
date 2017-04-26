package lib

import (
	"fmt"
	"regexp"
	"strings"
)

type GithubMarkdownRegexUpdater struct {
	OpenPrefix  string
	OpenSuffix  string
	ClosePrefix string
	CloseSuffix string
	Formatter   formatter
}

func urlFromID(id string) string {
	return fmt.Sprintf("https://%s", strings.Replace(id, ":", "/", -1))
}

func regexpForURL(url string) (*regexp.Regexp, error) {
	reg := fmt.Sprintf(`\(%s\)`, regexp.QuoteMeta(url))

	return regexp.Compile(reg)
}

func (u GithubMarkdownRegexUpdater) delimiters(stat Stats) (open, close string) {
	open = fmt.Sprintf("%s%s%s", u.OpenPrefix, stat.ID, u.OpenSuffix)
	close = fmt.Sprintf("%s%s%s", u.ClosePrefix, stat.ID, u.CloseSuffix)
	return
}

func (u GithubMarkdownRegexUpdater) Update(page string, stats []Stats) (updateResult, error) {
	result := updateResult{}

	for _, stat := range stats {
		formattedStats, err := u.Formatter.format(stat)
		if err != nil {
			return result, err
		}

		githubURL := urlFromID(stat.ID)

		r, err := regexpForURL(githubURL)
		if err != nil {
			return result, err
		}

		open, close := u.delimiters(stat)

		replacement := fmt.Sprintf("(%s) %s%s%s", githubURL, open, formattedStats, close)

		newPage := r.ReplaceAllString(page, replacement)

		if newPage != page {
			result.stats.added++
			result.stats.found++

			page = newPage
		}
	}

	result.page = page

	return result, nil
}
