package sy2dg

/*
Semester shows semester of lecture.
*/
type Semester int

const (
	Spring Semester = iota + 1
	SpringCenterrized
	Autumn
	AutumnCenterrized
	Centerrized
	ThroughYear
	One
	Two
	Three
	Four
)

/*
CompareResult represents the compare result of two semester objects.
*/
type CompareResult int

const (
	NotComparable CompareResult = iota + 1
	Before
	SameSemester
	After
)

func (cr CompareResult) String() string {
	switch cr {
	case NotComparable:
		return "NotComparable"
	case Before:
		return "Before"
	case After:
		return "After"
	case SameSemester:
		return "SameSemester"
	}
	return "Unknown"
}

/*
IsCenterrized returns the receiver semester object is centerrized.
*/
func (s Semester) IsCenterrized() bool {
	return s == SpringCenterrized || s == AutumnCenterrized || s == Centerrized
}

/*
CompareTo compares two Semester objects and returns the result.
*/
func (s Semester) CompareTo(other Semester) CompareResult {
	if s <= ThroughYear && other > ThroughYear ||
		s > ThroughYear && other <= ThroughYear {
		return NotComparable
	}
	if s.IsSameSemester(other) {
		return SameSemester
	}
	if s == ThroughYear || other == ThroughYear {
		return NotComparable
	}
	if s < other {
		return Before
	}
	return After
}

/*
IsSameSemester returns receiver semester object and given semester object are in the same sememester.
*/
func (s Semester) IsSameSemester(other Semester) bool {
	return s == other ||
		s == Spring && other == SpringCenterrized || s == SpringCenterrized && other == Spring ||
		s == Autumn && other == AutumnCenterrized || s == AutumnCenterrized && other == Autumn
}

/*
Available confirms the receiver semester object is availab.e
*/
func (s Semester) Available() bool {
	return s >= Spring && s <= Four
}

/*
ParseSemester parses given string and returns semester object.
*/
func ParseSemester(str string) Semester {
	switch str {
	case "春学期":
		return Spring
	case "秋学期":
		return Autumn
	case "春学期集中":
		return SpringCenterrized
	case "秋学期集中":
		return AutumnCenterrized
	case "集中":
		return Centerrized
	case "通年":
		return ThroughYear
	case "第1期":
		return One
	case "第2期":
		return Two
	case "第3期":
		return Three
	case "第4期":
		return Four
	}
	return Semester(0)
}

func (s Semester) String() string {
	switch s {
	case Spring:
		return "春学期"
	case Autumn:
		return "秋学期"
	case SpringCenterrized:
		return "春学期集中"
	case AutumnCenterrized:
		return "秋学期集中"
	case Centerrized:
		return "集中"
	case ThroughYear:
		return "通年"
	case One:
		return "第1期"
	case Two:
		return "第2期"
	case Three:
		return "第3期"
	case Four:
		return "第4期"
	default:
		return "不明"
	}
}
