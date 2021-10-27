package anagrams

import (
	"fmt"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func CheckAnagrams() {
	var s, j, res string
	res = "0"

	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &j)

	if s == j {
		res = "1"
		fmt.Print(res)
		return
	}

	s = strings.TrimSuffix(strings.ToLower(s), "\n")
	j = strings.TrimSuffix(strings.ToLower(j), "\n")

	if s == j {
		res = "1"
		fmt.Print(res)
		return
	}

	if len(s) == len(j) {
		sSort1 := SortString(s)
		sSort2 := SortString(j)

		if sSort1 == sSort2 {
			res = "1"
		}
	}

	fmt.Print(res)
}
