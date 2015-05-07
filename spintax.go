package spintax

import (
	"bytes"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var re = regexp.MustCompile("(?U)\\{.*\\}")

// Spin creates a spin from list of strings.
func Spin(strs []string) string {
	var buf bytes.Buffer
	buf.WriteString("{")
	buf.WriteString(strings.Join(strs, "|"))
	buf.WriteString("}")
	return buf.String()
}

// Unspin creates a string from given spin.
func Unspin(spin string) string {
	replace := func(match string) string {
		parts := strings.Split(match[1:len(match)-1], "|")
		res := parts[rand.Intn(len(parts))]
		return res
	}
	return re.ReplaceAllStringFunc(spin, replace)
}

// Count returns the number of variant of the spin.
func Count(spin string) int {
	count := 1
	matches := re.FindAllString(spin, -1)
	for _, match := range matches {
		parts := strings.Split(match[1:len(match)-1], "|")
		if len(parts) >= 1 {
			count *= len(parts)
		}
	}
	return count
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
