package sy2dg

import "testing"

func TestCompareResult(t *testing.T) {
	testdata := []struct {
		give CompareResult
		wont string
	}{
		{NotComparable, "NotComparable"},
		{Before, "Before"},
		{SameSemester, "SameSemester"},
		{After, "After"},
		{After + 1, "Unknown"},
	}
	for _, td := range testdata {
		if td.give.String() != td.wont {
			t.Errorf("CompareResult.String() did not match, wont %s, got %s", td.wont, td.give.String())
		}
	}
}

func TestIsCenterrized(t *testing.T) {
	testdata := []struct {
		give Semester
		wont bool
	}{
		{Spring, false},
		{SpringCenterrized, true},
		{Autumn, false},
		{AutumnCenterrized, true},
		{Centerrized, true},
		{ThroughYear, false},
		{One, false},
		{Two, false},
		{Three, false},
		{Four, false},
	}
	for _, td := range testdata {
		if td.give.IsCenterrized() != td.wont {
			t.Errorf("Semester.IsCenterrized did not match, wont %v, got %v", td.wont, td.give.IsCenterrized())
		}
	}
}
