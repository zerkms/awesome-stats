package github_repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerkms/awesome-stats/parser"
)

func TestURLParserSuccessful(t *testing.T) {
	url := "https://github.com/zerkms/foo"

	owner, project, err := splitOwnerAndProject(parser.URL(url))

	assert.Equal(t, "zerkms", owner)
	assert.Equal(t, "foo", project)
	assert.Nil(t, err)
}

func TestURLParserErr(t *testing.T) {
	url := "https://github.com/"

	_, _, err := splitOwnerAndProject(parser.URL(url))

	assert.NotNil(t, err)
	assert.Equal(t, "Github url in wrong format: https://github.com/", err.Error())
}
