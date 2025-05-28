package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ravsii/forcealias/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "tests")
	analysistest.Run(t, testdata, analyzer.Analyzer, "alias")
}
