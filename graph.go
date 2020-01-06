package sy2dg

import (
	"sort"
	"strconv"
	"strings"
)

type DataSet struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Weight int    `json:"weight"`
	Arrows string `json:"arrows"`
}
type Node struct {
	Name         string        `json:"name"`
	Group        string        `json:"group"`
	TeacherNames []string      `json:"teacherNames"`
	Semester     string        `json:"semester"`
	Weight       int           `json:"weight"`
	ID           string        `json:"id"`
	Value        *SyllabusData `json:"-"`
	URL          string        `json:"url"`
}

type syllabusName struct {
	id    string
	name  string
	index int
}

func sortAndUniq(slice []int) []int {
	sort.Slice(slice, func(v1, v2 int) bool {
		if slice[v1] < slice[v2] {
			return true
		}
		return false
	})
	results := []int{}
	for i, value := range slice {
		if i == 0 || results[len(results)-1] != value {
			results = append(results, value)
		}
	}
	return results
}

func (sd *SyllabusData) containsLectureName(name string) bool {
	return strings.Contains(sd.Outline, name) || strings.Contains(sd.SpecialNotes, name)
}

func findRelations(syllabus *SyllabusData, names []*syllabusName) []int {
	results := []int{}
	for _, nai := range names {
		if syllabus.containsLectureName(nai.name) {
			results = append(results, nai.index)
		}
	}
	return sortAndUniq(results)
}

func createEdge(s []*SyllabusData, si, ti int) Edge {
	if s[si].IsSameSemester(s[ti]) {
		if s[si].ID < s[ti].ID {
			return Edge{Source: s[si].ID, Target: s[ti].ID, Arrows: "to, from"}
		}
		return Edge{Source: s[ti].ID, Target: s[si].ID, Arrows: "to, from"}
	}
	if s[si].IsPreviousOf(s[ti]) && !s[si].IsAfterOf(s[ti]) {
		return Edge{Source: s[si].ID, Target: s[ti].ID, Arrows: "to"}
	}
	return Edge{Source: s[ti].ID, Target: s[ti].ID, Arrows: "to"}
}

func ContainsEdge(edges []Edge, edge Edge) bool {
	for _, e := range edges {
		if e.Source == edge.Source && e.Target == edge.Target {
			return true
		}
	}
	return false
}

func countUp(edge Edge, sources []Edge) int {
	count := 0
	for _, e := range sources {
		if e.Source == edge.Source && e.Target == edge.Target {
			count++
		}
	}
	return count
}

func countUpValue(results []Edge, sources []Edge) []Edge {
	for i, source := range results {
		count := countUp(source, sources)
		results[i].Weight = count
	}
	return results
}
func removeDuplicationAndSelfDirection(edges []Edge) []Edge {
	results := []Edge{}
	for _, edge := range edges {
		if edge.Source != edge.Target && !ContainsEdge(results, edge) {
			results = append(results, edge)
		}
	}
	return countUpValue(results, edges)
}

func createEdges(syllabuses []*SyllabusData, names []*syllabusName) []Edge {
	edges := []Edge{}
	for i, syllabus := range syllabuses {
		relations := findRelations(syllabus, names)
		for _, rel := range relations {
			edges = append(edges, createEdge(syllabuses, i, rel))
		}
	}
	return removeDuplicationAndSelfDirection(edges)
}

func findName(names []*syllabusName, name string) bool {
	for _, n := range names {
		if n.name == name {
			return true
		}
	}
	return false
}

func newSyllabusName(syllabus *SyllabusData, name string, index int) *syllabusName {
	return &syllabusName{id: syllabus.ID, name: name, index: index}
}

func extractLectureNamesAndTheirAliases(syllabuses []*SyllabusData) []*syllabusName {
	names := []*syllabusName{}
	for i, syllabus := range syllabuses {
		if !findName(names, syllabus.LectureName) {
			names = append(names, newSyllabusName(syllabus, syllabus.LectureName, i))
		}
		for _, alias := range syllabus.Aliases {
			if !findName(names, alias) {
				names = append(names, newSyllabusName(syllabus, alias, i))
			}
		}
	}
	return names
}

func newNode(syllabus *SyllabusData) *Node {
	return &Node{
		ID: syllabus.ID, Name: syllabus.LectureName, Group: strconv.Itoa(syllabus.Grade),
		Value: syllabus, URL: syllabus.URL, Semester: syllabus.Semester.String(),
		TeacherNames: syllabus.TeacherNames,
	}
}
func toNodeSlice(syllabuses []*SyllabusData) []Node {
	results := []Node{}
	for _, syllabus := range syllabuses {
		node := newNode(syllabus)
		results = append(results, *node)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].ID < results[j].ID
	})
	return results
}

func MakeGraph(syllabuses []*SyllabusData) *DataSet {
	names := extractLectureNamesAndTheirAliases(syllabuses)
	edges := createEdges(syllabuses, names)
	ds := &DataSet{Nodes: toNodeSlice(syllabuses), Edges: edges}
	// ds.updateWeight()
	return ds
}
