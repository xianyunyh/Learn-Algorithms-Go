package string

import (
	"log"
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	toLowerCase("Hello")
	strs := []string{"flower", "flow", "flight"}
	s := longestCommonPrefix(strs)
	log.Println(s)
	i := strStr("mississippi", "issip")
	log.Println(i)
	log.Println(strings.Index("mississippi", "issip"))

	log.Println(lengthOfLastWord("a "))
	detectCapitalUse("FlaG")
}
