package analyzer

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type imports map[string]string

var (
	flagAliases          imports
	flagIgnoreDot        bool
	flagIgnoreUnderscore bool
)

func (f *imports) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *imports) Set(value string) error {
	flagAliases = make(imports)

	for alias := range strings.SplitSeq(value, ",") {
		kv := strings.SplitN(alias, "=", 2)
		if len(kv) == 1 {
			return fmt.Errorf("bad alias: %s", alias)
		}

		k := strings.TrimSpace(kv[1])
		v := strings.TrimSpace(kv[0])

		(*f)[k] = v
	}

	return nil
}

var Analyzer = &analysis.Analyzer{
	Name:     "forcealias",
	Doc:      "forces aliases for given imports",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

const (
	flagName  = "force-alias"
	flagUsage = ` forces aliases for given pairs. It's a comma-separated list,
	where each item has is parsed as k=v. Example:
	newAlias=yourPackage,nurl=net/url`
)

func init() {
	Analyzer.Flags.Var(&flagAliases, flagName, flagUsage)
	Analyzer.Flags.BoolVar(&flagIgnoreDot, "ignore-dot", false, "ignores dot imports")
	Analyzer.Flags.BoolVar(&flagIgnoreUnderscore, "ignore-underscore", false, "ignores underscore imports")
}

func run(pass *analysis.Pass) (any, error) {
	if len(flagAliases) == 0 {
		return nil, fmt.Errorf("no aliases, use %s", flagName)
	}

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		importDecl, ok := node.(*ast.ImportSpec)
		if !ok {
			return
		}

		var gotAlias string
		if name := importDecl.Name; name != nil {
			gotAlias = name.Name
		}

		if flagIgnoreDot && gotAlias == "." {
			return
		}

		if flagIgnoreUnderscore && gotAlias == "_" {
			return
		}

		path := importDecl.Path.Value
		path = strings.Trim(path, `"\`)

		if wantAlias, ok := flagAliases[path]; ok && wantAlias != gotAlias {
			pass.Reportf(node.Pos(), "should be aliased as %s", wantAlias)
		}
	})

	return nil, nil
}
