package crawl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	in, out string
}

func TestNormalize(t *testing.T) {
	assert := assert.New(t)
	tests := []test{
		// Valid
		{in: "aaronoellis.com", out: "http://aaronoellis.com"},
		{in: "aaronoellis.com/", out: "http://aaronoellis.com"},
		{in: "http://aaronoellis.com", out: "http://aaronoellis.com"},
		{in: "aaronoellis.com/route", out: "http://aaronoellis.com/route"},

		// Invalid - should generate errors
		{in: "", out: ""},
	}
	for _, t := range tests {
		out, err := normalize(t.in)
		if t.out == "" {
			assert.NotNil(
				err, "An empty result should have been generated for %s", t.in,
			)
		} else {
			assert.Equal(t.out, out.String())
		}
	}
}
