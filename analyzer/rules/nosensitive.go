package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type NoSensitiveRule struct{}

var sensitiveKeywords = []string{
	"password",
	"passwd",
	"pwd",
	"api_key",
	"apikey",
	"token",
	"secret",
	"auth",
	"credential",
}

func (NoSensitiveRule) Check(pass *analysis.Pass, expr ast.Expr, value string, pos token.Pos) {

	switch v := expr.(type) {

	case *ast.BinaryExpr:
		pass.Report(analysis.Diagnostic{
			Pos:     v.Pos(),
			End:     v.End(),
			Message: "log message may contain sensitive data",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "remove string concatenation from log message",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     v.Pos(),
							End:     v.End(),
							NewText: []byte(strconv.Quote(value)),
						},
					},
				},
			},
		})
		return

	case *ast.Ident:
		pass.Reportf(v.Pos(), "log message may contain sensitive data")
		return

	case *ast.BasicLit:
		if v.Kind != token.STRING {
			return
		}

		value := strings.ToLower(strings.Trim(v.Value, "\"`"))

		for _, keyword := range sensitiveKeywords {
			if strings.Contains(value, keyword) {
				pass.Reportf(v.Pos(), "log message may contain sensitive data")
				return
			}
		}
	}
}
