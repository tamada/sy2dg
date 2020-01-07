package sy2dg

import "testing"

func TestBuildSyllabus(t *testing.T) {
	parser := NewJSONParser()
	builder := NewSyllabusBuilder(parser, "syllabus[0-9]\\.json", "http://localhost")
	syllabuses := builder.ReadSyllabuses("testdata")
	if len(syllabuses) != 2 {
		t.Errorf("syllabus data did not match, wont 2, got %d", len(syllabuses))
	}
	if syllabuses[0].LectureName != "ソフトウェア工学I" && syllabuses[0].URL != "http://localhost/syllabus1.json" {
		t.Errorf("syllabus data did not match, got %v", syllabuses[0])
	}
}
