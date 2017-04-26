package lib

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyFormatter struct{}

func (dummyFormatter) format(s Stats) (string, error) {
	return fmt.Sprintf("%s:%d", s.ID, s.Stars), nil
}

type errorFormatter struct{}

func (errorFormatter) format(s Stats) (string, error) {
	return "", errors.New("always")
}

func TestReplaceSingle(t *testing.T) {
	f := dummyFormatter{}
	m := MarkerUpdater{
		OpenPrefix:  "open-prefix ",
		OpenSuffix:  " open-suffix",
		ClosePrefix: "close-prefix ",
		CloseSuffix: " close-suffix",
		Formatter:   f,
	}

	s := Stats{
		ID:    "type-id",
		Stars: 42,
	}

	page := "some text before open-prefix type-id open-suffix10close-prefix type-id close-suffix some text after"

	actual, err := m.Update(page, []Stats{s})
	stats := actual.stats

	assert.Equal(t, "some text before open-prefix type-id open-suffixtype-id:42close-prefix type-id close-suffix some text after", actual.page)
	assert.Equal(t, 1, stats.updated)
	assert.Equal(t, 0, stats.added)
	assert.Equal(t, 1, stats.found)

	assert.Nil(t, err)
}

func TestReplaceMultiple(t *testing.T) {
	f := dummyFormatter{}
	m := MarkerUpdater{
		OpenPrefix:  "op ",
		OpenSuffix:  " os",
		ClosePrefix: "cp ",
		CloseSuffix: " cs",
		Formatter:   f,
	}

	s1 := Stats{
		ID:    "ID",
		Stars: 42,
	}

	s2 := Stats{
		ID:    "id2",
		Stars: 123,
	}

	page := "some text before op ID oscp ID cs text between op ID os old value cp ID cs text after op id2 oscp id2 cs some more text"

	actual, err := m.Update(page, []Stats{s1, s2})
	stats := actual.stats

	assert.Equal(t, "some text before op ID osID:42cp ID cs text between op ID osID:42cp ID cs text after op id2 osid2:123cp id2 cs some more text", actual.page)
	assert.Equal(t, 2, stats.updated)
	assert.Equal(t, 0, stats.added)
	assert.Equal(t, 2, stats.found)

	assert.Nil(t, err)
}

func TestReplaceErrorFormatter(t *testing.T) {
	f := errorFormatter{}
	m := MarkerUpdater{
		Formatter: f,
	}
	s := Stats{}

	_, err := m.Update("", []Stats{s})

	assert.NotNil(t, err)
}
