package bencode

import (
	"testing"
)


func it(t *testing.T, in string, exp int64, exp_err bool) {
	d := NewDecoder([]byte(in))
	i, err := d.Decode()

	if !exp_err {
		if err != nil {
			t.Errorf("got error %s instead of %d", err.String(), exp)
		}
		if i != exp {
			t.Errorf("expected %d, got %d\n", exp, i)
		}
	} else {
		if err == nil {
			t.Errorf("[in: %s], [res: %d], expected error. got result",
				in,
				i)
		}
	}
}

func TestInteger(t *testing.T) {
	it(t, "i23e", 23, false)
	it(t, "i124145124e", 124145124, false)
	it(t, "i15155", 0, true)
	it(t, "55", 55, true)
}
