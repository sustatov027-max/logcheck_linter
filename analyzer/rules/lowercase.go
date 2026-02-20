package rules

import (
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type LowercaseRule struct{}

func (LowercaseRule) Check(pass *analysis.Pass, value string, pos token.Pos) {
	if value == "" {
		return
	}

	first := []rune(value)[0]

	if unicode.IsUpper(first) {
		pass.Reportf(pos, "log message should start with lowercase letter")
	}
}