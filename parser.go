package sy2dg

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

/*
JSONParser parses syllabus data from formatted in JSON.
*/
type JSONParser struct {
}

/*
NewJSONParser creates JsonParser and returns it.
*/
func NewJSONParser() *JSONParser {
	return new(JSONParser)
}

/*
Parse method parses given reader as json format of syllabus and build SyllabusData.
This method returns built SyllabusData and error if it occured.
*/
func (parser *JSONParser) Parse(reader io.Reader, fileName string) (*SyllabusData, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	data := SyllabusData{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	data.Semester = ParseSemester(data.SemesterString)

	return &data, nil
}
