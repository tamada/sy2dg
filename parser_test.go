package sy2dg

import (
	"os"
	"testing"
)

func TestJSONLoad(t *testing.T) {
	parser := NewJSONParser()
	file, _ := os.Open("testdata/syllabus1.json")
	defer file.Close()
	sd, err := parser.Parse(file, "testdata/syllabus1.json")
	if err != nil {
		t.Errorf(err.Error())
	}
	if sd.LectureName != "ソフトウェア工学I" || sd.Credit != 2 ||
		len(sd.TeacherNames) != 1 || sd.TeacherNames[0] != "玉田　春昭" ||
		sd.Grade != 2 || sd.Semester != Autumn {
		t.Errorf("parse error of \"testdata/syllabus1.json\"")
	}
}

func TestJSONReadError(t *testing.T) {
	parser := NewJSONParser()
	file, _ := os.Open("testdata/invalid.json")
	defer file.Close()
	_, err := parser.Parse(file, "testdata/invalid.json")
	if err == nil {
		t.Errorf("no error by reading invalid json file.")
	}
}
