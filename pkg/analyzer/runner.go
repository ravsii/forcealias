package analyzer

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Aliases in a [packagePath] => expectedAlias style.
//
// Example:
//
//	Aliases: map[string]string{
//		"net/url": "defaultNetUrl"
//	}
type Aliases map[string]string

type Config struct {
	// Import aliases. See [Aliases] comment for examples.
	Aliases Aliases

	// If true, dot imports will not be reported as invalid.
	IgnoreDot bool

	// If true, underscore imports will not be reported as invalid.
	IgnoreUnderscore bool
}

type runner struct {
	Config Config
}

func (f *Aliases) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *Aliases) Set(value string) error {
	*f = make(Aliases)

	for alias := range strings.SplitSeq(value, ",") {
		kv := strings.SplitN(alias, "=", 2)
		if len(kv) == 1 {
			return fmt.Errorf("bad alias: %s", alias)
		}

		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])

		(*f)[k] = v
	}

	return nil
}

func (r *runner) run(pass *analysis.Pass) (any, error) {
	if len(r.Config.Aliases) == 0 {
		return nil, nil
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

		if r.Config.IgnoreDot && gotAlias == "." {
			return
		}

		if r.Config.IgnoreUnderscore && gotAlias == "_" {
			return
		}

		path := importDecl.Path.Value
		path = strings.Trim(path, `"\`)

		if wantAlias, ok := r.Config.Aliases[path]; ok && wantAlias != gotAlias {
			pass.Reportf(node.Pos(), "should be aliased as %s", wantAlias)
		}
	})

	return nil, nil
}
