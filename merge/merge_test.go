package merge_test

import (
	"testing"

	"github.com/davelondon/jennifer/merge"
)

var cases = []tc{
	{
		dsc: `very simple example`,
		old: `package foo`,
		new: `package foo`,
		exp: `package foo`,
	},
}

func TestMerge(t *testing.T) {
	for i, c := range cases {
		out := merge.Merge([]byte(c.old), []byte(c.new))
		if string(out) != c.exp {
			t.Errorf("Test case %d failed. Description: %s\nExpected:\n%s\nOutput:\n%s", i, c.dsc, c.exp, out)
		}
	}
}

// a test case
type tc struct {
	// description for locating the test case
	dsc string
	// old source
	old string
	// new generated source
	new string
	// expected merged source
	exp string
}
