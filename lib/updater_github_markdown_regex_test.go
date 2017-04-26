package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type parenthesesFormatter struct{}

func (parenthesesFormatter) format(s Stats) (string, error) {
	return fmt.Sprintf("(%d)", s.Stars), nil
}

func TestReplaceMarkdownSimple(t *testing.T) {
	page := "[name](https://github.com/owner/project) some additional text"
	s := Stats{
		ID:    "github.com:owner:project",
		Stars: 42,
	}

	expected := "[name](https://github.com/owner/project) <!-- github.com:owner:project -->(42)<!-- /github.com:owner:project --> some additional text"

	u := GithubMarkdownRegexUpdater{
		OpenPrefix:  "<!-- ",
		OpenSuffix:  " -->",
		ClosePrefix: "<!-- /",
		CloseSuffix: " -->",
		Formatter:   parenthesesFormatter{},
	}

	actual, err := u.Update(page, []Stats{s})
	stats := actual.stats

	assert.Nil(t, err)

	assert.Equal(t, expected, actual.page)
	assert.Equal(t, stats.added, 1)
	assert.Equal(t, stats.updated, 0)
	assert.Equal(t, stats.found, 1)
}
