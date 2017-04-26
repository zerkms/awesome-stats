package lib

type URL string

type URLExtractor interface {
	Extract(page string) []URL
}

type Parser struct {
	Extractors map[string]URLExtractor
}

func (p Parser) Parse(page string) map[string][]URL {
	result := make(map[string][]URL)

	for id, extractor := range p.Extractors {
		result[id] = extractor.Extract(page)
	}

	return result
}
