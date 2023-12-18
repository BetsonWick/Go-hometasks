package _defer

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_FirstSequence(t *testing.T) {
	var builder strings.Builder
	PrintSequence1(&builder)

	assert.Equal(t, "135642", builder.String())
}

func Test_SecondSequence(t *testing.T) {
	var builder strings.Builder
	PrintSequence2(&builder)

	assert.Equal(t, "1345287960", builder.String())
}
