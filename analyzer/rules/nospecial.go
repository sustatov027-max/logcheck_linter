package rules

import (
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

var allowedRegexp = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9\s]+$`)

type NoSpeacialRule struct{}

func (NoSpeacialRule) Check(pass *analysis.Pass, value string, pos token.Pos){
	if !allowedRegexp.MatchString(value){
		pass.Reportf(pos, "log message contains forbidden characters or emoji")
	}
}