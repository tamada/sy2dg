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

func TestCompareSemester(t *testing.T) {
	testdata := []struct {
		receiver  Semester
		parameter Semester
		wont      CompareResult
	}{
		{Spring, Autumn, Before},
		{Autumn, Spring, After},
		{Spring, SpringCenterrized, SameSemester},
		{AutumnCenterrized, Autumn, SameSemester},
		{Autumn, ThroughYear, NotComparable},
		{Spring, Spring, SameSemester},
		{Autumn, One, NotComparable},
	}
	for _, td := range testdata {
		result := td.receiver.CompareTo(td.parameter)
		if result != td.wont {
			t.Errorf("%s.CompareTo(%s) did not match, wont %s, got %v", td.receiver, td.parameter, td.wont, result)
		}
	}
}

func TestParseAndSpring(t *testing.T) {
	testdata := []struct {
		give string
		wont Semester
	}{
		{"春学期", Spring},
		{"春学期集中", SpringCenterrized},
		{"秋学期", Autumn},
		{"秋学期集中", AutumnCenterrized},
		{"通年", ThroughYear},
		{"集中", Centerrized},
		{"第1期", One},
		{"第2期", Two},
		{"第3期", Three},
		{"第4期", Four},
		{"不明", 0},
	}
	for _, td := range testdata {
		got := ParseSemester(td.give)
		if got != td.wont {
			t.Errorf("ParseSemester(%s) did not match, wont %s, got %s", td.give, td.wont, got)
		}
		got2 := td.wont.String()
		if got2 != td.give {
			t.Errorf("%s.String() did not match, wont %s, got %s", td.wont, td.give, got2)
		}
	}
}

func TestIsCenterrized(t *testing.T) {
	testdata := []struct {
		give          Semester
		wont          bool
		availableFlag bool
	}{
		{Spring, false, true},
		{SpringCenterrized, true, true},
		{Autumn, false, true},
		{AutumnCenterrized, true, true},
		{Centerrized, true, true},
		{ThroughYear, false, true},
		{One, false, true},
		{Two, false, true},
		{Three, false, true},
		{Four, false, true},
		{Four + 1, false, false},
	}
	for _, td := range testdata {
		if td.give.IsCenterrized() != td.wont {
			t.Errorf("Semester.IsCenterrized did not match, wont %v, got %v", td.wont, td.give.IsCenterrized())
		}
		if td.give.Available() != td.availableFlag {
			t.Errorf("%s.Available did not match, wont %v, got %v", td.give, td.availableFlag, !td.availableFlag)
		}
	}
}
