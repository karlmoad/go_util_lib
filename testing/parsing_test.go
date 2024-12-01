package testing

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/parsing"
	"os"
	"testing"
)

func TestParsing_parser(t *testing.T) {

	bites, _ := os.ReadFile("test.ebnf")
	par, err := parsing.NewParserForDialect(parsing.GRAMMAR)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Starting Parse.")
	expressions, err := par.Parse(string(bites))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(fmt.Sprintf("Number of expressions parsed: %d", len(expressions)))
}
