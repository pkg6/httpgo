package t

import (
	"strconv"
	"strings"
)

// IntExplode
// 1,2,3=>[]int{1,2,3}
func IntExplode(s, sep string) ([]int, error) {
	var is []int
	for _, sp := range strings.Split(s, sep) {
		i, err := strconv.Atoi(sp)
		if err != nil {
			return is, err
		}
		is = append(is, i)
	}
	return is, nil
}

// IntImplode
//[]int{1,2,3}=>1,2,3
func IntImplode(is []int, sep string) string {
	ss := make([]string, len(is))
	for i, num := range is {
		ss[i] = strconv.Itoa(num)
	}
	return strings.Join(ss, sep)
}
