package shellwords

import (
	"regexp"
	"strings"
)

var specialChars = regexp.MustCompile("([^A-Za-z0-9_\\-.,:\\/@\n])")
var lf = regexp.MustCompile("\\n")

func Escape(s string) string {
	if len(s) == 0 {
		return "''"
	}

	s = specialChars.ReplaceAllString(s, "\\$1")
	s = lf.ReplaceAllString(s, "'\\n'")

	return s
}

func Join(ss []string) string {
	escaped := make([]string, len(ss))
	for i, s := range ss {
		escaped[i] = Escape(s)
	}
	return strings.Join(escaped, " ")
}
