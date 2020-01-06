package sy2dg

import (
	"reflect"
	"testing"
)

func TestCreatingGraph(t *testing.T) {
	builder := NewSyllabusBuilder(NewJSONParser(), `[0-9]+\.json`, "https:/syllabus.kyoto-su.ac.jp/syllabus/html/2019")
	syllabuses := builder.ReadSyllabuses("testdata/json")
	ds := NewDataSet(syllabuses)
	if len(ds.Nodes) != 5 {
		t.Errorf("node length did not match, wont %d, got %d (syllabus length: %d)", 10, len(ds.Nodes), len(syllabuses))
	}
	wontRelations := []Edge{
		Edge{Source: "3978", Target: "3966"}, // 発プロ -> ソフトウェア工学I
		Edge{Source: "3978", Target: "3967"}, // 発プロ -> ソフトウェア工学II
		Edge{Source: "3966", Target: "3967"}, // ソフトウェア工学I -> ソフトウェア工学II
		Edge{Source: "3967", Target: "3869"}, // ソフトウェア工学II -> オペレーティングシステム
		Edge{Source: "3956", Target: "3978"}, // 情報化社会論 -> 発プロ
		Edge{Source: "3956", Target: "3966"}, // 情報化社会論 -> ソフトウェア工学I
	}
	if len(ds.Edges) != len(wontRelations) {
		t.Errorf("edge length did not match, wont %d, got %d", len(wontRelations), len(ds.Edges))
	}
	for _, r := range wontRelations {
		if !findEdge(ds, r) {
			t.Errorf("edge did not found %v", r)
		}
	}
}

func findEdge(node *DataSet, edge Edge) bool {
	for _, e := range node.Edges {
		if e.Target == edge.Target && e.Source == edge.Source {
			return true
		}
	}
	return false
}

func TestSortAndUniq(t *testing.T) {
	testdata := []struct {
		give []int
		wont []int
	}{
		{[]int{1, 2, 5, 3, 2, 1}, []int{1, 2, 3, 5}},
	}
	for _, td := range testdata {
		got := sortAndUniq(td.give)
		if !reflect.DeepEqual(got, td.wont) {
			t.Errorf("the result of sortAndUniq(%v) did not match, wont %v, got %v", td.give, td.wont, got)
		}
	}
}
