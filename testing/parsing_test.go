package testing

import (
	"github.com/karlmoad/go_util_lib/parsing"
	"github.com/karlmoad/go_util_lib/parsing/dialects/antlrG4"
	"github.com/karlmoad/go_util_lib/parsing/dialects/grammar"
	"os"
	"testing"
)

func TestGrammar_parser(t *testing.T) {
	bites, _ := os.ReadFile("test.txt")

	dialect := grammar.NewGrammarDialect()
	par, err := parsing.NewParserForDialect(dialect)
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

	count := 0
	for _, expression := range expressions {
		if expression.Elem().Kind() == grammar.RULE_ELEM {
			count++
		}
	}

	if count != 4 {
		t.Errorf("Expected 4 expressions of type RULE_ELEM, got %d", count)
	}

}

func TestAntlrG4_parser(t *testing.T) {
	bites, _ := os.ReadFile("test2.txt")

	//src := string(bites)
	//t.Error(src[855:871])

	dialect := antlrG4.NewGrammarDialect()
	par, err := parsing.NewParserForDialect(dialect)
	if err != nil {
		t.Error(err)
	}

	expressions, err := par.Parse(string(bites))
	if err != nil {
		t.Error(err)
	}

	if len(expressions) != 8 {
		t.Errorf("Expected 4 expressions, got %d", len(expressions))
	}

	count := 0
	for _, expression := range expressions {
		if expression.Elem().Kind() == grammar.RULE_ELEM {
			count++
		}
	}

	if count != 8 {
		t.Errorf("Expected 4 expressions of type RULE_ELEM, got %d", count)
	}
}
