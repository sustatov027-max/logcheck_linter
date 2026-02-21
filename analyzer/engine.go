package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Check(pass *analysis.Pass, expr ast.Expr, literal string, pos token.Pos)
}

var rulesSlice []Rule

func RegisterRule(r Rule) {
	rulesSlice = append(rulesSlice, r)
}