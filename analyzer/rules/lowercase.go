package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type LowercaseRule struct{}

func (LowercaseRule) Check(pass *analysis.Pass, expr ast.Expr, value string, pos token.Pos) {
	if value == "" {
		return
	}

	runes := []rune(value)
	if len(runes) == 0 {
		return
	}

	if !unicode.IsUpper(runes[0]) {
		return
	}

	if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.STRING {
		runes[0] = unicode.ToLower(runes[0])
		fixed := strconv.Quote(string(runes))

		pass.Report(analysis.Diagnostic{
			Pos:     lit.Pos(),
			End:     lit.End(),
			Message: "log message should start with lowercase letter",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "convert first letter to lowercase",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(fixed),
						},
					},
				},
			},
		})
		return
	}
	pass.Reportf(pos, "log message should start with lowercase letter")
}
