package str

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func Slice(input string, split string) []string {
	var result []string
	for _, s := range strings.Split(input, split) {
		s = strings.TrimSpace(s)
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

func StringOrDefault(input string, defaultValue string) string {
	if input == "" {
		return defaultValue
	} else {
		return input
	}
}

func FormattedStrOrEmpty(format, input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf(format, input)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func Random(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func LongestCommonPrefix(input []string) string {
	var longestPrefix string = ""
	var endPrefix = false
	if len(input) > 0 {
		sort.Strings(input)
		first := string(input[0])
		last := string(input[len(input)-1])
		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}
