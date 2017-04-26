package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const envName = "AWESOME_STATS_GITHUB_TOKEN"
const defaultTemplate = "(★{{.Stars}})"

type config struct {
	in          string
	out         string
	template    string
	githubToken string
}

func parseConfig() (config, error) {
	c := config{}

	flag.StringVar(&c.in, "in", "", "the input file")
	flag.StringVar(&c.out, "out", "", "the path to write output result")
	flag.StringVar(&c.template, "template", defaultTemplate, "stats template")

	flag.Parse()

	return validateConfig(c)
}

func validateConfig(c config) (config, error) {
	if c.in == "" {
		return c, errors.New("input file parameter is required")
	}

	if c.out == "" {
		c.out = c.in
	}

	githubToken := os.Getenv(envName)
	if githubToken == "" {
		return c, fmt.Errorf("Please set a github token as `%s` environment variable", envName)
	}
	c.githubToken = githubToken

	return c, nil
}
