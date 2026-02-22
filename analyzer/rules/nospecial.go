package rules

import (
	"go/ast"
	"go/token"
	"regexp"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

var allowedRegexp = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9\s]+$`)
var forbiddenCharsRegexp = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9\s]+`)

type NoSpeacialRule struct{}

func (NoSpeacialRule) Check(pass *analysis.Pass, expr ast.Expr, value string, pos token.Pos) {
	if !allowedRegexp.MatchString(value) {
		fixed := forbiddenCharsRegexp.ReplaceAllString(value, "")
		pass.Report(analysis.Diagnostic{
			Pos:     pos,
			End:     expr.End(),
			Message: "log message contains forbidden characters or emoji",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "remove special characters from log message",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     pos,
							End:     expr.End(),
							NewText: []byte(strconv.Quote(fixed)),
						},
					},
				},
			},
		})
	}
}