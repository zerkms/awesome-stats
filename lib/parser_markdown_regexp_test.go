package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	page := `
* [flac](https://github.com/eaburns/flac) - A native Go FLAC decoder.
* [flac](https://github.com/mewkiz/flac) - A native Go FLAC decoder.
* [vorbis](https://github.com/mccoyst/vorbis) - A "native" Go Vorbis decoder (uses CGO, but has no dependencies).`

	p := MarkdownRegexParser{}
	r := p.Extract(page)

	expected := []URL{
		"https://github.com/eaburns/flac",
		"https://github.com/mewkiz/flac",
		"https://github.com/mccoyst/vorbis",
	}

	assert.Equal(t, expected, r)
}
