package sy2dg

import (
	"reflect"
	"testing"
)

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
