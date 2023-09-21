package utils

import (
	"regexp"
	"strings"
)

var (
	spaceReg = regexp.MustCompile(`\s+`)
)

func RemoveDuplicatedSpace(in string) (out string) {
	out = spaceReg.ReplaceAllString(in, " ")
	out = strings.TrimSpace(out)
	return
}
