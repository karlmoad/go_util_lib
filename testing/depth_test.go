package testing

import (
	"github.com/karlmoad/go_util_lib/common/state"
	"testing"
)

func TestDepth_State(t *testing.T) {
	var depth state.Depth

	if depth.CurrentDepth() != 0 && depth.CurrentState() != false {
		t.Errorf("Current depth should be 0 [state: %t], got %d [%t]", false, depth.CurrentDepth(), depth.CurrentState())
	}
	depth.Increase()
	if depth.CurrentDepth() != 1 && depth.CurrentState() != true {
		t.Errorf("Current depth should be 1 [state: %t] after increase, got %d [%t]", true, depth.CurrentDepth(), depth.CurrentState())
	}
	depth.Increase()
	if depth.CurrentDepth() != 2 && depth.CurrentState() != true {
		t.Errorf("Current depth should be 2 [state: %t] after increase, got %d [%t]", true, depth.CurrentDepth(), depth.CurrentState())
	}
	depth.Increase()
	if depth.CurrentDepth() != 3 && depth.CurrentState() != true {
		t.Errorf("Current depth should be 3 [state: %t] after increase, got %d [%t]", true, depth.CurrentDepth(), depth.CurrentState())
	}
	depth.Decrease()
	if depth.CurrentDepth() != 2 && depth.CurrentState() != true {
		t.Errorf("Current depth should be 2 [state: %t] after decrease, got %d [%t]", true, depth.CurrentDepth(), depth.CurrentState())
	}
	depth.Decrease()
	if depth.CurrentDepth() != 1 && depth.CurrentState() != true {
		t.Errorf("Current depth should be 1 [state: %t] after decrease, got %d [%t]", true, depth.CurrentDepth(), depth.CurrentState())
	}
	depth.Decrease()
	if depth.CurrentDepth() != 0 && depth.CurrentState() != false {
		t.Errorf("Current depth should be 0 [state: %t] after decrease, got %d [%t]", false, depth.CurrentDepth(), depth.CurrentState())
	}
}
