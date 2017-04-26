package lib

import (
	"bytes"
	"text/template"
)

type GoTemplateFormatter struct {
	StatsFormat string
}

func (f GoTemplateFormatter) format(s Stats) (string, error) {
	t := template.New("StatsFormat")
	t, err := t.Parse(f.StatsFormat)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	if err := t.Execute(&result, s); err != nil {
		return "", err
	}

	return result.String(), nil
}
