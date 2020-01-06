package sy2dg

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type SyllabusBuilder struct {
	Parser  Parser
	Pattern *regexp.Regexp
	BaseURL string
}

func NewSyllabusBuilder(parser Parser, patternString, baseURL string) *SyllabusBuilder {
	pattern := regexp.MustCompile(patternString)
	return &SyllabusBuilder{Pattern: pattern, Parser: parser, BaseURL: baseURL}
}

func (ts *SyllabusBuilder) IsTarget(str string) bool {
	return ts.Pattern.Match([]byte(str))
}

func (sb *SyllabusBuilder) buildURL(data *SyllabusData, targetPath string) {
	if data.URL == "" && sb.BaseURL != "" {
		data.URL = path.Join(sb.BaseURL, filepath.Base(targetPath))
	}
}

func readSyllabus(targetPath string, sb *SyllabusBuilder) (*SyllabusData, error) {
	file, err := os.Open(targetPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	data, err := sb.Parser.Parse(file, targetPath)
	if err == nil {
		sb.buildURL(data, targetPath)
	}
	return data, nil
}

func convertToSlice(s map[string]*SyllabusData) []*SyllabusData {
	results := []*SyllabusData{}
	for _, v := range s {
		results = append(results, v)
	}
	sort.Slice(results, func(i, j int) bool {
		return strings.Compare(results[i].URL, results[j].URL) < 0
	})
	return results
}

func (sb *SyllabusBuilder) ReadSyllabuses(dir string) []*SyllabusData {
	syllabuses := map[string]*SyllabusData{}
	filepath.Walk(dir, func(targetPath string, info os.FileInfo, err error) error {
		fileName := filepath.Base(targetPath)
		if sb.IsTarget(fileName) {
			data, err := readSyllabus(targetPath, sb)
			if err == nil {
				syllabuses[fileName] = data
			}
		}
		return err
	})
	slice := convertToSlice(syllabuses)
	fmt.Fprintf(os.Stderr, "read %d syllabuses\n", len(slice))
	return slice
}
