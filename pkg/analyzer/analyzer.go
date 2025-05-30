package analyzer

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const (
	flagForceAliasName       = "force-alias"
	flagIgnoreDotName        = "ignore-dot"
	flagIgnoreUnderscoreName = "ignore-underscore"

	flagForceAliasUsage = `forces aliases for given pairs. It's a comma-separated list,` +
		`where each item has is parsed as k=v. Example:` +
		`oldName=newName,net/url=defaultUrl`
)

// NewAnalyzer creates a new analyzer that parses flags for runner config.
func NewAnalyzer() *analysis.Analyzer {
	var r runner

	a := &analysis.Analyzer{
		Name:     "forcealias",
		Doc:      "forces aliases for given imports",
		Run:      r.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.Var(&r.Config.Aliases, flagForceAliasName, flagForceAliasUsage)
	a.Flags.BoolVar(&r.Config.IgnoreDot, flagIgnoreDotName, false, "ignores dot imports")
	a.Flags.BoolVar(&r.Config.IgnoreUnderscore, flagIgnoreUnderscoreName, false, "ignores underscore imports")

	return a
}

// NewAnalyzer creates a new analyzer that parses flags for runner config.
func NewAnalyzerWithConfig(c Config) *analysis.Analyzer {
	r := runner{Config: c}

	a := &analysis.Analyzer{
		Name:     "forcealias",
		Doc:      "forces aliases for given imports",
		Run:      r.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	return a
}
