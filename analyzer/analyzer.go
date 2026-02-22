package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/sustatov027-max/logcheck_linter/analyzer/rules"
)

var Analyzer = &analysis.Analyzer{
	Name: "logrules",
	Doc:  "checks log messages (custom linter)",
	Run:  run,
}

func init() {
	RegisterRule(rules.LowercaseRule{})
	RegisterRule(rules.EnglishRule{})
	RegisterRule(rules.NoSpeacialRule{})
	RegisterRule(rules.NoSensitiveRule{})
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isSupportedLoggerCall(pass, call) {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			value := extractString(call.Args[0])
			if value == "" {
				return true
			}

			for _, r := range rulesSlice {
				r.Check(pass, call.Args[0], value, call.Args[0].Pos())
			}

			return true
		})
	}

	return nil, nil
}

func isSupportedLoggerCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if !isLogMethod(selector.Sel.Name) {
		return false
	}

	if ident, ok := selector.X.(*ast.Ident); ok {
		if obj, ok := pass.TypesInfo.Uses[ident]; ok {
			if pkgName, ok := obj.(*types.PkgName); ok {
				if pkgName.Imported().Path() == "log/slog" {
					return true
				}
			}
		}
	}

	typ := pass.TypesInfo.TypeOf(selector.X)
	if typ == nil {
		return false
	}

	ptr, ok := typ.(*types.Pointer)
	if ok {
		typ = ptr.Elem()
	}

	named, ok := typ.(*types.Named)
	if !ok {
		return false
	}

	if named.Obj() == nil || named.Obj().Pkg() == nil {
		return false
	}

	pkgPath := named.Obj().Pkg().Path()

	if pkgPath == "log/slog" {
		return true
	}

	if pkgPath == "go.uber.org/zap" {
		return true
	}

	return false
}

func isLogMethod(name string) bool {
	switch name {
	case "Info", "Error", "Warn", "Debug", "Fatal":
		return true
	default:
		return false
	}
}

func extractString(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.BasicLit:
		if v.Kind == token.STRING {
			return strings.Trim(v.Value, "\"`")
		}
	case *ast.BinaryExpr:
		left := extractString(v.X)
		right := extractString(v.Y)
		return left + right
	}
	return ""
}
