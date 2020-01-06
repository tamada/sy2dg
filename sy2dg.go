package sy2dg

import (
	"io"
	"strings"
)

const Version = "1.0.0"

type Parser interface {
	Parse(reader io.Reader, fileName string) (*SyllabusData, error)
}

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

func (sd *SyllabusData) IsSameSemester(other *SyllabusData) bool {
	return sd.Grade == other.Grade && sd.Semester.CompareTo(other.Semester) == SameSemester
}

func (sd *SyllabusData) IsPreviousOf(other *SyllabusData) bool {
	if sd.Grade == other.Grade {
		result := sd.Semester.CompareTo(other.Semester)
		if result == SameSemester {
			return sd.ID < other.ID
		}
		return result == Before
	}
	return sd.Grade < other.Grade
}

func (sd *SyllabusData) IsAfterOf(other *SyllabusData) bool {
	if sd.Grade == other.Grade {
		result := sd.Semester.CompareTo(other.Semester)
		if result == SameSemester {
			return sd.ID > other.ID
		}
		return result == After
	}
	return sd.Grade > other.Grade
}
