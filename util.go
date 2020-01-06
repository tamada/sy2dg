package sy2dg

import "path/filepath"

/*
Uniq removes duplicated entries from given slice and returns resultant slice.
*/
func Uniq(slice []string) []string {
	results := []string{}
	for _, item := range slice {
		if !Contains(results, item) {
			results = append(results, item)
		}
	}
	return results
}

/*
RemoveItems removes given items from given slice and returns resultant slice.
*/
func RemoveItems(slice []string, items []string) []string {
	results := []string{}
	for _, item := range items {
		results = removeItem(slice, item)
	}
	return results
}

func removeItem(slice []string, item string) []string {
	results := []string{}
	for _, sliceItem := range slice {
		if sliceItem != item {
			results = append(results, sliceItem)
		}
	}
	return results
}

/*
FileNameWithoutExt returns the fileName only without extention.
Example:
     FileNameWithoutExt("1234.json")          // => 1234
	 FileNameWithoutExt("/path/of/5678.json") // => 5678
*/
func FileNameWithoutExt(fileName string) string {
	name := filepath.Base(fileName)
	ext := filepath.Ext(fileName)
	return name[:len(name)-len(ext)]
}

/*
Contains examins inclusion of given item in the slice.
If slice contains given item, it returns true, otherwise false.
*/
func Contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}
	return false
}

/*
AppendAsSet appends the item to the slice if slice did not have item.
If slice has item in the entries, it does nothing.
*/
func AppendAsSet(slice []string, item string) []string {
	if !Contains(slice, item) {
		slice = append(slice, item)
	}
	return slice
}
