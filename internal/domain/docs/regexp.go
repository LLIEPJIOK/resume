package docs

import "regexp"

var expRuRegex = regexp.MustCompile(`Опыт разработки:\s*(\d+)\s*лет`)

func getExperienceRegexp() *regexp.Regexp {
	return expRuRegex
}
