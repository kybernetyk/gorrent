package bencode

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	//first an int64
	var o1 int64
	err := Unmarshal([]byte("i9223372036854775807e"), &o1)
	if err != nil {
		t.Errorf("Got error while unmarshaling 'i23e': %s", err.String())
	} else if o1 != 9223372036854775807 {
		t.Errorf("'i23e' did not unmarshal into 9223372036854775807: %d", o1)
	}

}

func TestUnmarshalString(t *testing.T) {
	type stringTest struct {
		input     string
		output    string
		expecterr bool
	}

	tests := []stringTest{
		{"4:test", "test", false},
		{"10:abcdefghij", "abcdefghij", false},
		{"4:l", "", true},
		{"2134lol", "", true},
		{"4:longtestlol", "long", false},
	}

	for _, test := range tests {
		var o string
		err := unmarshalString([]byte(test.input), &o)
		if !test.expecterr && o != test.output {
			t.Errorf("input was %s, expected %s, result %s",
				test.input,
				test.output,
				o)
		}

		if test.expecterr && err == nil {
			t.Errorf("[i: %s], [ex: %s], [res: %s], expected error but got none!",
				test.input,
				test.output,
				o)
		}
		if !test.expecterr && err != nil {
			t.Errorf("[i: %s], [ex: %s], [res: %s], expected no error. but got one: %s",
				test.input,
				test.output,
				o,
				err.String())
		}
	}
}

func TestUnmarshalInt64(t *testing.T) {
	type integerTest struct {
		input     string
		output    int64
		expecterr bool
	}

	tests := []integerTest{
		{"i23e", 23, false},
		{"i0e", 0, false},
		{"i46e", 46, false},
		{"afafaf", 0, true},
		{"i23343", 0, true},
		{"i9223372036854775807e", 9223372036854775807, false},
		//		{"i18446744073709551615e", 18446744073709551615, false}, //in case we do uint64
	}

	for _, test := range tests {
		var o int64
		err := unmarshalInt64([]byte(test.input), &o)
		if !test.expecterr && o != test.output {
			t.Errorf("input was %s, expected %d, result %d", test.input, test.output, o)
		}

		if test.expecterr && err == nil {
			t.Errorf("[i: %s], [ex: %d], [res: %d], expected error. but got none!",
				test.input,
				test.output,
				o)
		}

		if !test.expecterr && err != nil {
			t.Errorf("[i: %s], [ex: %d], [res: %d], expected no error. but got one: %s",
				test.input,
				test.output,
				o,
				err.String())
		}
	}
}
