package lib

import "regexp"

type MarkdownRegexParser struct{}

var urlRegex = regexp.MustCompile(`\[[^\]]+\]\((https://github.com/[^/]+/[^/?#)]+)\)`)

func (MarkdownRegexParser) Extract(page string) []URL {
	result := []URL{}

	for _, match := range urlRegex.FindAllStringSubmatch(page, -1) {
		result = append(result, URL(match[1]))
	}

	return result
}
