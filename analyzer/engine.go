package analyzer

import (
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Check(pass *analysis.Pass, literal string, pos token.Pos)
}

var rulesSlice []Rule

func RegisterRule(r Rule) {
	rulesSlice = append(rulesSlice, r)
}