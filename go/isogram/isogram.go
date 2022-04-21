package isogram

import (
	"sort"
	"strings"
)

func IsIsogram(word string) bool {
	if word == "" {
		return true
	}

	replacer := strings.NewReplacer("-", "", " ", "")
	str := replacer.Replace(word)
	str = strings.ToLower(str)

	strArr := strings.Split(str, "")
	sort.Strings(strArr)

	for i := 0; i < len(strArr)-1; i++ {
		if strArr[i] == strArr[i+1] {
			return false
		}
	}

	return true
}
