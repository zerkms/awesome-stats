package shared

import (
	"github.com/zerkms/awesome-stats/lib"
	"github.com/zerkms/awesome-stats/lib/github_repository"
)

type Builder struct {
	githubToken string
	template    string
}

func (b *Builder) SetTemplate(template string) {
	b.template = template
}

func (b *Builder) SetGithubToken(token string) {
	b.githubToken = token
}

func (b *Builder) Build() lib.AwesomeStats {
	p := lib.Parser{
		Extractors: map[string]lib.URLExtractor{
			"github.com": lib.MarkdownRegexParser{},
		},
	}

	githubRepository := github_repository.GithubRepository{
		AccessToken: b.githubToken,
	}

	repositories := map[string]lib.StatsCollectionRepository{
		"github.com": lib.SequentialRepository{
			Repository: githubRepository,
		},
	}

	formatter := lib.GoTemplateFormatter{
		StatsFormat: b.template,
	}

	ghRegexUpdater := lib.GithubMarkdownRegexUpdater{
		OpenPrefix:  "<!-- ",
		OpenSuffix:  " -->",
		ClosePrefix: "<!-- /",
		CloseSuffix: " -->",
		Formatter:   formatter,
	}

	markerUpdater := lib.MarkerUpdater{
		OpenPrefix:  "<!-- ",
		OpenSuffix:  " -->",
		ClosePrefix: "<!-- /",
		CloseSuffix: " -->",
		Formatter:   formatter,
	}

	updater := lib.NewCompositeUpdater([]lib.Updater{markerUpdater, ghRegexUpdater})

	awesomeStats := lib.NewAwesomeStats(p, repositories, updater)

	return awesomeStats
}
