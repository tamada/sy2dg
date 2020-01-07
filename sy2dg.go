package sy2dg

import (
	"io"
	"strings"
)

/*
Version specifies the version of the sy2dg.
*/
const Version = "1.0.0"

/*
Parser is an interface for reading syllabus data and builds an instance of the syllabus.
*/
type Parser interface {
	Parse(reader io.Reader, fileName string) (*SyllabusData, error)
}

/*
SyllabusData represents the syllabus.
*/
type SyllabusData struct {
	ID             string   `json:"ID"`
	LectureName    string   `json:"lecture-name"`
	Aliases        []string `json:"aliases"`
	TeacherNames   []string `json:"teacher-names"`
	Grade          int      `json:"grade"`
	Semester       Semester `json:"-"`
	Credit         int      `json:"credit"`
	Outline        string   `json:"outline"`
	SpecialNotes   string   `json:"special-notes"`
	SemesterString string   `json:"semester"`
	URL            string   `json:"url"`
}

func (sd *SyllabusData) containsLectureName(name string) bool {
	return strings.Contains(sd.Outline, name) || strings.Contains(sd.SpecialNotes, name)
}

/*
IsSameSemester returns true if lectures of receiver and argument
conduct in the same semester and for the same grades.
*/
func (sd *SyllabusData) IsSameSemester(other *SyllabusData) bool {
	return sd.Grade == other.Grade && sd.Semester.CompareTo(other.Semester) == SameSemester
}

/*
IsPreviousOf returns true if receiver lecture conduct before argument lecture.
Examples:
    * { grade: 2, semester: Spring }.IsPreviousOf({ grade: 3, semester: Spring}) => true
	* { grade: 2, semester: Spring }.IsPreviousOf({ grade: 2, semester: Autumn}) => true
	* { grade: 2, semester: Spring }.IsPreviousOf({ grade: 2, semester: Spring}) => false
	* { grade: 3, semester: Spring }.IsPreviousOf({ grade: 2, semester: Spring}) => false
	* { grade: 2, semester: Autumn }.IsPreviousOf({ grade: 2, semester: Spring}) => false
*/
func (sd *SyllabusData) IsPreviousOf(other *SyllabusData) bool {
	if sd.Grade == other.Grade {
		result := sd.Semester.CompareTo(other.Semester)
		return result == Before
	}
	return sd.Grade < other.Grade
}

/*
IsAfterOf returns true if receiver lecture conduct behinde of argument lecture.
Examples:
    * { grade: 2, semester: Spring }.IsAfterOf({ grade: 3, semester: Spring}) => false
	* { grade: 2, semester: Spring }.IsAfterOf({ grade: 2, semester: Autumn}) => false
	* { grade: 2, semester: Spring }.IsAfterOf({ grade: 2, semester: Spring}) => false
	* { grade: 3, semester: Spring }.IsAfterOf({ grade: 2, semester: Spring}) => true
	* { grade: 2, semester: Autumn }.IsAfterOf({ grade: 2, semester: Spring}) => true
*/
func (sd *SyllabusData) IsAfterOf(other *SyllabusData) bool {
	if sd.Grade == other.Grade {
		result := sd.Semester.CompareTo(other.Semester)
		return result == After
	}
	return sd.Grade > other.Grade
}
