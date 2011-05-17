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
		bytesread	int64
	}

	tests := []stringTest{
		{"4:test", "test", false, 6},
		{"10:abcdefghij", "abcdefghij", false, 13},
		{"4:l", "", true, 0},
		{"2134lol", "", true, 0},
		{"4:longtestlol", "long", false, 6},
	}

	for _, test := range tests {
		var o string
		br, err := unmarshalString([]byte(test.input), &o)
		if !test.expecterr && o != test.output {
			t.Errorf("input was %s, expected %s, result %s",
				test.input,
				test.output,
				o)
		}

		if !test.expecterr && br != test.bytesread {
			t.Errorf("[i: %s], [ex: %s], [res: %s], expected br of %d but got %d!",
				test.input,
				test.output,
				o,
				test.bytesread,
				br)
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
		bytesread int64
	}

	tests := []integerTest{
		{"i23e", 23, false, 4},
		{"i0e", 0, false, 3},
		{"i46e", 46, false, 4},
		{"afafaf", 0, true, 0},
		{"i23343", 0, true, 0},
		{"i9223372036854775807e", 9223372036854775807, false, int64(len("i9223372036854775807e"))},
		//		{"i18446744073709551615e", 18446744073709551615, false}, //in case we do uint64
	}

	for _, test := range tests {
		var o int64
		br, err := unmarshalInt64([]byte(test.input), &o)
		if !test.expecterr && o != test.output {
			t.Errorf("input was %s, expected %d, result %d", test.input, test.output, o)
		}

		if !test.expecterr && br != test.bytesread {
			t.Errorf("[i: %s], [ex: %d], [res: %d], expected br of %d but got %d!",
				test.input,
				test.output,
				o,
				test.bytesread,
				br)
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
