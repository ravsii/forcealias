package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ravsii/forcealias/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "tests")

	tests := []struct {
		packageName string
		Config      analyzer.Config
	}{
		{
			packageName: "no_alias", Config: analyzer.Config{},
		},
		{
			packageName: "alias", Config: analyzer.Config{
				Aliases: analyzer.Aliases{
					"fmt":           "testAlias",
					"net/url":       "notDot",
					"encoding/json": "notUnderscore",
				},
			},
		},
		{
			packageName: "dot_ignore",
			Config: analyzer.Config{
				Aliases:   analyzer.Aliases{"io": "shouldBeIgnored"},
				IgnoreDot: true,
			},
		},
		{
			packageName: "dot_ignore_alias",
			Config: analyzer.Config{
				Aliases: analyzer.Aliases{
					"fmt": "testAlias",
					"io":  "shouldBeIgnored",
				},
				IgnoreDot: true,
			},
		},
		{
			packageName: "underscore_ignore",
			Config:      analyzer.Config{IgnoreUnderscore: true},
		},
		{
			packageName: "underscore_ignore_alias",
			Config: analyzer.Config{
				Aliases:          analyzer.Aliases{"fmt": "testAlias"},
				IgnoreUnderscore: true,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.packageName, func(t *testing.T) {
			t.Parallel()

			a := analyzer.NewAnalyzerWithConfig(tc.Config)
			analysistest.Run(t, testdata, a, tc.packageName)
		})
	}
}
