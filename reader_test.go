package csvutils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {

	readTests := []struct {
		name  string
		input string
		comma rune
		want  []map[string]string
		error string
	}{
		{
			name:  "simple file",
			input: "a,b,c\n1,2,3",
			want:  []map[string]string{{"a": "1", "b": "2", "c": "3"}},
			comma: ',',
			error: "",
		},
		{
			name:  "empty file",
			input: "",
			want:  nil,
			comma: ',',
			error: "",
		},
		{
			name:  "| delimiter",
			input: "A|B|C\n1|2|3",
			comma: '|',
			want: []map[string]string{
				{"A": "1", "B": "2", "C": "3"},
			},
			error: "",
		},
		{
			name:  "Duplicate Headers",
			input: "A,B,C,A\n1,2,3,4",
			comma: ',',
			want:  nil,
			error: "Multiple indices with the same name 'A'",
		},
		{
			name:  "long row",
			input: "A,B,C\n1,2,3,4",
			comma: ',',
			want:  nil,
			error: "record on line 2: wrong number of fields",
		},
		{
			name:  "short row",
			input: "A,B,C\n1,2",
			comma: ',',
			want:  nil,
			error: "record on line 2: wrong number of fields",
		},
	}

	for _, tt := range readTests {
		t.Run(tt.name, func(t *testing.T) {
			reader := NewReader(strings.NewReader(tt.input))
			reader.Reader.Comma = tt.comma
			reader.Columns, _ = reader.ReadHeader()

			got, err := reader.ReadAll()

			assert.Equal(t, got, tt.want)

			if tt.error != "" {
				assert.Equal(t, err.Error(), tt.error)
			}
		})

	}
}

func TestReadHeaders(t *testing.T) {

	input := strings.NewReader("A,B,C\n1,2,3")
	r := NewReader(input)

	headers, err := r.ReadHeader()
	exp := []string{"A", "B", "C"}
	if !reflect.DeepEqual(headers, exp) {
		t.Errorf("out=%q, want=%q", headers, exp)
	}

	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
}
