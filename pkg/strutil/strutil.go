package strutil

import "strings"

func SplitAnyCleanPrefix(s, seps, prefix string) []string {
	trimmed := strings.TrimPrefix(s, prefix)
	return SplitAnyClean(strings.TrimSpace(trimmed), seps)
}

func SplitAnyClean(s, seps string) []string {
	f := func(c rune) bool {
		return strings.ContainsRune(seps, c)
	}
	l := strings.FieldsFunc(s, f)
	cleared := make([]string, 0, len(l))

	for _, item := range l {
		item = strings.TrimSpace(item)
		if item != "" {
			cleared = append(cleared, item)
		}
	}

	return cleared
}
