package matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyMatcher(t *testing.T) {
	m := NewRegexPatternMatcher()
	assert.False(t, m.Matches("/test"))
	assert.False(t, m.Matches("*"))
}
func TestAllMatcher(t *testing.T) {
	m := NewRegexPatternMatcher("*")
	assert.True(t, m.Matches("/test"))
}
func TestCustomMatcher(t *testing.T) {
	m := NewRegexPatternMatcher("*/service/test")
	assert.False(t, m.Matches("/test"))
	assert.True(t, m.Matches("/some/service/test"))
}
