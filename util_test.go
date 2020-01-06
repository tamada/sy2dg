package sy2dg

import (
	"reflect"
	"testing"
)

func IsEqualSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func TestRemoveItems(t *testing.T) {
	testdata := []struct {
		slice           []string
		removeFromSlice []string
		wontSlice       []string
	}{
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, []string{"b", "c"}, []string{"a", "d", "e", "f", "g"}},
	}
	for _, td := range testdata {
		gotSlice := RemoveItems(td.slice, td.removeFromSlice)
		if IsEqualSlice(td.wontSlice, gotSlice) {
			t.Errorf("the result of RemoveItems did not match, wont %v, got %v", td.wontSlice, gotSlice)
		}
	}
}

func TestFileNameWithoutExt(t *testing.T) {
	testdata := []struct {
		giveName string
		wontName string
	}{
		{"/path/of/1234.json", "1234"},
		{"5678.json", "5678"},
		{"5678.sample.html", "5678.sample"},
	}

	for _, td := range testdata {
		result := FileNameWithoutExt(td.giveName)
		if result != td.wontName {
			t.Errorf("the result of FileNameWithoutExt did not match, wont %s, got %s", td.wontName, result)
		}
	}
}

func TestContains(t *testing.T) {
	testdata := []struct {
		slice    []string
		item     string
		wontFlag bool
	}{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "d", false},
	}
	for _, td := range testdata {
		if Contains(td.slice, td.item) != td.wontFlag {
			t.Errorf("Contains(%v, %s) failed, wont %v, got %v", td.slice, td.item, td.wontFlag, !td.wontFlag)
		}
	}
}

func TestAddAsSet(t *testing.T) {
	testdata := []struct {
		slice      []string
		addedItem  string
		wontLength int
	}{
		{[]string{"a", "b", "c"}, "a", 3},
		{[]string{"a", "b", "c"}, "b", 3},
		{[]string{"a", "b", "c"}, "d", 4},
	}
	for _, td := range testdata {
		slice := AppendAsSet(td.slice, td.addedItem)
		if len(slice) != td.wontLength {
			t.Errorf("AddAsSet(%v, %s) failed, length wont %v, got %v", td.slice, td.addedItem, td.wontLength, len(slice))
		}
		if len(slice) != td.wontLength && slice[len(slice)-1] != td.addedItem {
			t.Errorf("added item was not match, wont %s, got %s", td.addedItem, slice[len(slice)-1])
		}
	}
}

func TestRefineAsSet(t *testing.T) {
	testdata := []struct {
		slice     []string
		wontSlice []string
	}{
		{[]string{"a", "b", "c", "a", "b"}, []string{"a", "b", "c"}},
		{[]string{"a", "a", "b", "a", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c", "d", "e"}},
	}
	for _, td := range testdata {
		slice := Uniq(td.slice)
		if len(slice) != len(td.wontSlice) || !reflect.DeepEqual(slice, td.wontSlice) {
			t.Errorf("RefineAsSet(%v) failed, wont %v, got %v", td.slice, td.wontSlice, slice)
		}
	}
}
