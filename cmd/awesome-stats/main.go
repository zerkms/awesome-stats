package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zerkms/awesome-stats/lib/shared"
)

func main() {
	c, err := parseConfig()
	handleError(err)

	page, err := ioutil.ReadFile(c.in)
	handleError(err)

	builder := shared.Builder{}
	builder.SetGithubToken(c.githubToken)
	builder.SetTemplate(c.template)
	awesomeStats := builder.Build()

	r, err := awesomeStats.Update(string(page))
	handleError(err)

	if err := writeResult(c.out, r); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeResult(filename, result string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.WriteString(result)
	return err
}
