package docs

import "regexp"

var expRuRegex = regexp.MustCompile(`Опыт разработки.*(\d+)\+?\s*(лет|года|год)`)

func getExperienceRegexp() *regexp.Regexp {
	return expRuRegex
}
