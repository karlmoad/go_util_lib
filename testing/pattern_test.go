package testing

import (
	"github.com/karlmoad/go_util_lib/common/regex"
	"strings"
	"testing"
)

func TestPattern_MatchSourceStart(t *testing.T) {

	testString := "[17] | Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ultricies."

	p1 := regex.NewPattern(`^\[[0-9]+/]\s*\|\s*`)
	p2 := regex.NewPattern(`^Lorem\s+`)
	p3 := regex.NewPattern(`ipsum\s+sit`)

	if v1, valid := p1.MatchSourceStart(testString); valid {
		testVal := "[17] | "
		if strings.Compare(v1, testVal) != 0 {
			t.Errorf("Failed to match string expected match:%t, value:%s got %s|%t", true, testVal, v1, valid)
		}
	}

	if v2, valid := p2.MatchSourceStart(testString); valid {
		t.Errorf("Expected no match, got %s|%t", v2, valid)
	}

	if v3, valid := p3.MatchSourceStart(testString); valid {
		t.Errorf("Expected no match, got %s|%t", v3, valid)
	}

}
