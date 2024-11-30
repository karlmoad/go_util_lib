package testing

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/parsing/dialect/grammar"
	"github.com/karlmoad/go_util_lib/parsing/parser"
	"os"
	"testing"
)

func TestParsing_parser(t *testing.T) {

	bites, _ := os.ReadFile("test.ebnf")
	dialect := grammar.NewGrammarDialect()

	p := parser.NewParser(string(bites), dialect)
	fmt.Println("Starting Parse.")
	expressions, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(fmt.Sprintf("Number of expressions parsed: %d", len(expressions)))
}
