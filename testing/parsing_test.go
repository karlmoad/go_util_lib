package testing

import (
	"github.com/karlmoad/go_util_lib/parsing"
	"os"
	"testing"
)

func TestParsing_parser(t *testing.T) {
	bites, _ := os.ReadFile("test.txt")
	par, err := parsing.NewParserForDialect(parsing.GRAMMAR)
	if err != nil {
		t.Error(err)
	}

	expressions, err := par.Parse(string(bites))
	if err != nil {
		t.Error(err)
	}

	if len(expressions) != 4 {
		t.Errorf("Expected 4 expressions, got %d", len(expressions))
	}
}
