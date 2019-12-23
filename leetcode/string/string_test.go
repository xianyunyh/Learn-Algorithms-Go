package string

import "testing"

import "log"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	s := longestCommonPrefix(strs)
	log.Println(s)
}
