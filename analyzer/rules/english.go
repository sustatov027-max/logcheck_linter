package rules

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type EnglishRule struct{}

func (EnglishRule) Check(pass *analysis.Pass, expr ast.Expr, value string, pos token.Pos) {
	if value == "" {
		return
	}

	if forbiddenCharsRegexp.MatchString(value) {
		return
	}

	if !isEnglishOnly(value) {
		if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			translateStr := Translate(value, "ru", "en")
			translateRune := strings.ToLower(translateStr)

			fixed := strconv.Quote(string(translateRune))

			pass.Report(analysis.Diagnostic{
				Pos:     lit.Pos(),
				End:     lit.End(),
				Message: "log message should contain only english characters",
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "convert message to english language",
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
		pass.Reportf(pos, "log message should contain only english characters")
	}
}

func Translate(text, from, to string) string {
	url := fmt.Sprintf("https://api.mymemory.translated.net/get?q=%s&langpair=%s|%s",
		url.QueryEscape(text), from, to)

	resp, err := http.Get(url)
	if err != nil {
		return "ошибка запроса"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	response := string(body)
	start := strings.Index(response, `"translatedText":"`)
	if start == -1 {
		return "не удалось найти перевод"
	}
	start += len(`"translatedText":"`)
	end := strings.Index(response[start:], `"`)
	if end == -1 {
		return "ошибка парсинга"
	}

	return response[start : start+end]
}

func isEnglishOnly(s string) bool {
	if s == "" {
		return true
	}
	
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == ' ' {
			continue
		}

		switch r {
		case '.', ',', '!', '?', ':', ';', '-', '\'', '"', '(', ')', '[', ']', '{', '}':
			continue
		}
		
		return false
	}
	return true
}