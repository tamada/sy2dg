package sy2dg

import "io"

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
