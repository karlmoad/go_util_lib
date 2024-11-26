package regex

import "regexp"

type Pattern struct {
	pattern *regexp.Regexp
}

func NewPattern(pattern string) *Pattern {
	return &Pattern{regexp.MustCompile(pattern)}
}

func (p *Pattern) MatchSourceStart(source string) (string, bool) {
	match := p.pattern.FindStringIndex(source)
	if match != nil && match[0] == 0 {
		return source[match[0]:match[1]], true
	} else {
		return "", false
	}
}
