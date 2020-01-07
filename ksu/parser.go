package ksu

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/tamada/sy2dg"
	"golang.org/x/text/unicode/norm"
	"gopkg.in/xmlpath.v1"
)

/*
HTMLParser parses syllabuses of KSU.
*/
type HTMLParser struct {
}

/*
NewHTMLParser generates an instance of sy2dg.Parser for parsing syllabuses of KSU formatted in HTML.
*/
func NewHTMLParser() *HTMLParser {
	return new(HTMLParser)
}

/*
Parse parses syllabus and returns generated SyllabusData.
*/
func (parser *HTMLParser) Parse(reader io.Reader, fileName string) (*sy2dg.SyllabusData, error) {
	root, err := xmlpath.ParseHTML(reader)
	if err != nil {
		return nil, err
	}
	data, err := buildData(root)
	if err != nil {
		return nil, err
	}
	data.ID = sy2dg.FileNameWithoutExt(fileName)
	updateAliases(data)
	return data, nil
}

var matcher = regexp.MustCompile("<([a-z],?)+>")

func updateAliases(data *sy2dg.SyllabusData) {
	nameBytes := []byte(data.LectureName)
	if matcher.Match(nameBytes) {
		stripName := matcher.ReplaceAll(nameBytes, []byte{})
		data.Aliases = append(data.Aliases, string(stripName))
	}
}

func buildData(root *xmlpath.Node) (*sy2dg.SyllabusData, error) {
	syllabus := new(sy2dg.SyllabusData)
	updateData(syllabus, root)
	updateAlias(syllabus)
	return syllabus, nil
}

func assignContent(assignTo *string, root *xmlpath.Node, xpathString string) {
	path := xmlpath.MustCompile(xpathString)
	value, ok := path.String(root)
	if ok {
		*assignTo = normalize(value)
	}
}

func assignContents(assignToArray []*string, root *xmlpath.Node, xpathString string) {
	path := xmlpath.MustCompile(xpathString)
	items := []string{}
	iter := path.Iter(root)
	for iter.Next() {
		node := iter.Node()
		items = append(items, node.String())
	}
	*assignToArray[0] = items[1]
	*assignToArray[1] = items[6]
}

var aliasRegexp = regexp.MustCompile(`\<([a-d],?)+\>`)

func updateAlias(data *sy2dg.SyllabusData) {
	if aliasRegexp.MatchString(data.LectureName) {
		alias := aliasRegexp.ReplaceAllString(data.LectureName, "")
		data.Aliases = append(data.Aliases, strings.TrimSpace(alias))
	}
}

func updateData(syllabus *sy2dg.SyllabusData, root *xmlpath.Node) {
	assignContent(&syllabus.LectureName, root, `//td[@class="syllabus_item_left syllabus_frame_TRB"]/span`)
	assignContents([]*string{&syllabus.Outline, &syllabus.SpecialNotes}, root, `//td[@class="syllabus_item_left syllabus_frame_LRB space_top_bottom"]/*`)
	updateByIterator(syllabus, root)
}

func updateByIterator(syllabus *sy2dg.SyllabusData, root *xmlpath.Node) {
	path := xmlpath.MustCompile(`//td[@class="syllabus_item_left syllabus_frame_RB"]/*`)
	items := []string{}
	iter := path.Iter(root)
	for iter.Next() {
		node := iter.Node()
		items = append(items, node.String())
	}
	syllabus.Semester = sy2dg.ParseSemester(items[0])
	syllabus.Grade = findGrade(items[2])
	syllabus.Credit = findCredit(items[3])
	syllabus.TeacherNames = findTeacherNames(items[4])
	syllabus.SemesterString = string(syllabus.Semester.String())
}

func findTeacherNames(names string) []string {
	return strings.Split(names, ",")
}

func findCredit(str string) int {
	credit := strings.TrimRight(str, "単位")
	return convertToNumber(normalize(credit))
}

func findGrade(str string) int {
	grade := strings.TrimRight(str, "年次")
	return convertToNumber(normalize(grade))
}

func normalize(str string) string {
	return string(norm.NFKC.Bytes([]byte(str)))
}

func convertToNumber(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return value
}
