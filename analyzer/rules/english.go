package rules

import (
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type EnglishRule struct{}

func (EnglishRule) Check(pass *analysis.Pass, value string, pos token.Pos) {
	for _, r := range value {
		if r > unicode.MaxASCII {
			pass.Reportf(pos, "log message should contain only english characters")
			return
		}
	}
}