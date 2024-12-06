package testing

import (
	"github.com/karlmoad/go_util_lib/generics/utils"
	"regexp"
	"testing"
)

func TestCollectionUtil_All(t *testing.T) {

	t1 := []bool{true, false, true}
	t2 := []bool{true, true, true}
	t3 := []bool{false, false, false}
	t4 := []int{2, 3, 4, 5, 6}
	t5 := []string{"FooBar", "FooBin", "FooBox"}

	p1 := func(i int) bool {
		return i > 1
	}

	regx := regexp.MustCompile("^FooB[a-z]+")

	p2 := func(s string) bool {
		return regx.MatchString(s)
	}

	if utils.All(t1, utils.True) {
		t.Error("Test 1 (True) expected result false got true")
	}

	if utils.All(t1, utils.False) {
		t.Error("Test 1 (False) expected result false got true")
	}

	if !utils.All(t2, utils.True) {
		t.Error("Test 2 expected result true got false")
	}

	if !utils.All(t3, utils.False) {
		t.Error("Test 3 expected result true got false")
	}

	if !utils.All(t4, p1) {
		t.Error("Test 4 expected result true got false")
	}

	if !utils.All(t5, p2) {
		t.Error("Test 5 expected result true got false")
	}
}

func TestCollectionUtil_Any(t *testing.T) {

	t1 := []bool{true, false, true}
	t2 := []int{2, -4, 1, -10, 6}
	t3 := []int{-2, -4, 1, -10, -6}
	t4 := []string{"FooBar", "F00Bin", "FoBox"}
	t5 := []string{"FaaBar", "F00Bin", "FoBox"}

	p1 := func(i int) bool {
		return i > 1
	}

	regx := regexp.MustCompile("^FooB[a-z]+")

	p2 := func(s string) bool {
		return regx.MatchString(s)
	}

	if !utils.Any(t1, utils.True) {
		t.Error("Test 1 (True) expected result true got false")
	}

	if !utils.Any(t1, utils.False) {
		t.Error("Test 1 (False) expected result true got false")
	}

	if !utils.Any(t2, p1) {
		t.Error("Test 2 expected result true got false")
	}

	if utils.Any(t3, p1) {
		t.Error("Test 3 expected result false got true")
	}

	if !utils.Any(t4, p2) {
		t.Error("Test 4 expected result true got false")
	}

	if utils.Any(t5, p2) {
		t.Error("Test 5 expected result false got true")
	}
}

func TestCollectionUtil_One(t *testing.T) {

	t1 := []bool{true, false, true}
	t2 := []int{2, -4, 1, -10, 6}
	t3 := []int{2, -4, 1, -10, -6}
	t4 := []string{"FooBar", "F00Bin", "FoBox"}
	t5 := []string{"FaaBar", "F00Bin", "FoBox"}

	p1 := func(i int) bool {
		return i > 1
	}

	regx := regexp.MustCompile("^FooB[a-z]+")

	p2 := func(s string) bool {
		return regx.MatchString(s)
	}

	if utils.One(t1, utils.True) {
		t.Error("Test 1 (True) expected result false got true")
	}

	if !utils.One(t1, utils.False) {
		t.Error("Test 1 (False) expected result true got false")
	}

	if utils.One(t2, p1) {
		t.Error("Test 2 expected result false got true")
	}

	if !utils.One(t3, p1) {
		t.Error("Test 3 expected result true got false")
	}

	if !utils.One(t4, p2) {
		t.Error("Test 4 expected result true got false")
	}

	if utils.One(t5, p2) {
		t.Error("Test 5 expected result false got true")
	}

}

func TestCollectionUtil_AtLeast(t *testing.T) {
	t1 := []bool{true, false, true, false, true, true, false}
	t2 := []int{2, -4, 1, -10, 6, -3}
	t3 := []int{2, -4, 1, -10, -6, -3}
	t4 := []string{"FooBar", "F00Bin", "FooBox"}
	t5 := []string{"FaaBar", "F00Bin", "FoBox"}

	p1 := func(i int) bool {
		return i > 1
	}

	regx := regexp.MustCompile("^FooB[a-z]+")

	p2 := func(s string) bool {
		return regx.MatchString(s)
	}

	if !utils.AtLeast(t1, 2, utils.True) {
		t.Error("Test 1 (True) expected result true got false")
	}

	if !utils.AtLeast(t1, 2, utils.False) {
		t.Error("Test 1 (False) expected result true got false")
	}

	if utils.AtLeast(t1, 5, utils.True) {
		t.Error("Test 1 (True) [neg] expected result false got true")
	}

	if utils.AtLeast(t1, 4, utils.False) {
		t.Error("Test 1 (False) [neg] expected result false got true")
	}

	if !utils.AtLeast(t2, 2, p1) {
		t.Error("Test 2 expected result true got false")
	}

	if utils.AtLeast(t3, 2, p1) {
		t.Error("Test 3 expected result false got true")
	}

	if !utils.AtLeast(t4, 2, p2) {
		t.Error("Test 4 expected result true got false")
	}

	if utils.AtLeast(t5, 2, p2) {
		t.Error("Test 5 expected result false got true")
	}
}

func TestCollectionUtil_N(t *testing.T) {
	t1 := []bool{true, false, true, false, true, true, false}
	t2 := []int{2, -4, 1, -10, 6, -3}
	t3 := []int{2, -4, 1, -10, -6, -3}
	t4 := []string{"FooBar", "F00Bin", "FooBox"}
	t5 := []string{"FaaBar", "F00Bin", "FoBox"}

	p1 := func(i int) bool {
		return i > 1
	}

	regx := regexp.MustCompile("^FooB[a-z]+")

	p2 := func(s string) bool {
		return regx.MatchString(s)
	}

	if utils.N(t1, 2, utils.True) {
		t.Error("Test 1 (True) expected result false got true")
	}

	if !utils.N(t1, 3, utils.False) {
		t.Error("Test 1 (False) expected result true got false")
	}

	if !utils.N(t2, 2, p1) {
		t.Error("Test 2 expected result true got false")
	}

	if utils.N(t3, 2, p1) {
		t.Error("Test 3 expected result false got true")
	}

	if !utils.N(t4, 2, p2) {
		t.Error("Test 4 expected result true got false")
	}

	if utils.N(t5, 2, p2) {
		t.Error("Test 5 expected result false got true")
	}
}

func TestCollectionUtil_Compare(t *testing.T) {
	t1 := []int{1, 2, 3, 4, 5}
	t2 := []int{1, 2, 3, 4, 5}
	t3 := []int{5, 4, 3, 2}
	t4 := []bool{true, true, false}
	t5 := []bool{true, true, false}
	t6 := []bool{true, false, false}
	t7 := []string{"One", "Two", "Three"}
	t8 := []string{"One", "Two", "Three"}
	t9 := []string{"One", "One", "Three"}

	if !utils.Compare(t1, t2) {
		t.Error("Boolean test 1 expected result true got false")
	}

	if utils.Compare(t1, t3) {
		t.Error("Boolean test 2 expected result false got true")
	}

	if !utils.Compare(t4, t5) {
		t.Error("Integer test 1 expected result true got false")
	}

	if utils.Compare(t4, t6) {
		t.Error("Integer test 2 expected result false got true")
	}

	if !utils.Compare(t7, t8) {
		t.Error("String test 1 expected result true got false")
	}

	if utils.Compare(t7, t9) {
		t.Error("String test 2 expected result false got true")
	}
}

func TestCollectionUtil_Map(t *testing.T) {

	t1 := []bool{true, false, true}
	t2 := []int{1, 2, 3, 4, 5}
	t3 := []string{"Foo", "Too", "Woo"}

	t1_e := []bool{false, true, false}
	t2_e := []int{2, 4, 6, 8, 10}
	t3_e := []string{"FooBar", "TooBar", "WooBar"}

	p1 := func(i bool) (bool, error) {
		return !i, nil
	}

	p2 := func(i int) (int, error) {
		return i * 2, nil
	}

	p3 := func(s string) (string, error) {
		return s + "Bar", nil
	}

	if rez, err := utils.Map(t1, p1); err == nil {
		if !utils.Compare(rez, t1_e) {
			t.Error("Map test 1 result not equal to expected")
		}
	} else {
		t.Error(err)
	}

	if rez, err := utils.Map(t2, p2); err == nil {
		if !utils.Compare(rez, t2_e) {
			t.Error("Map test 2 result not equal to expected")
		}
	} else {
		t.Error(err)
	}

	if rez, err := utils.Map(t3, p3); err == nil {
		if !utils.Compare(rez, t3_e) {
			t.Error("Map test 3 result not equal to expected")
		}
	} else {
		t.Error(err)
	}
}
