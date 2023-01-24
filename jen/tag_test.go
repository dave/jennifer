package jen_test

import (
	"bytes"
	"os"
	"testing"

	. "github.com/dave/jennifer/jen"
)

func TestTag_Render(t *testing.T) {
	type args struct {
		stmt *Statement
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Tag",
			args: args{
				stmt: Type().Id("foo").Struct(
					Id("Bar").String().Tag(map[string]string{"json": "bar"}),
				),
			},
			want: "./example/tag_001.go",
		},
		{
			name: "Tag with escaped characters",
			args: args{
				stmt: Type().Id("foo").Struct(
					Id("Bar").String().Tag(map[string]string{"valid": "matches(^(\\d*\\.)?\\d+$)"}),
				),
			},
			want: "./example/tag_002.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := NewFile("")

			expect, err := os.ReadFile(tt.want)
			if err != nil {
				t.Fatal("unable to read example file")
			}

			var got bytes.Buffer

			err = tt.args.stmt.RenderWithFile(&got, file)
			if err != nil {
				t.Fatal(err)
			}

			if got.String() != string(expect) {
				t.Fatalf("Got: %v, expect: %v", got.String(), expect)
			}
		})
	}
}
