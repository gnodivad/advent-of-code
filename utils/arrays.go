package utils

import "reflect"

func Contains(haystack interface{}, needle interface{}) bool {
	haystackValue := reflect.ValueOf(haystack)

	for i := 0; i < haystackValue.Len(); i++ {
		if haystackValue.Index(i).Interface() == needle {
			return true
		}
	}

	return false
}

// Based on Simple function in https://github.com/juliangruber/go-intersect/blob/master/intersect.go
func Intersect(a []string, b []string) []string {
	set := make([]string, 0)

	for i := 0; i < len(a); i++ {
		element := a[i]
		if Contains(b, element) {
			set = append(set, element)
		}
	}

	return set
}

func RemoveItem(haystack []string, needle string) []string {
	var newitems []string

	for _, i := range haystack {
		if i != needle {
			newitems = append(newitems, i)
		}
	}

	return newitems
}
