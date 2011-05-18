package main

import (
	"fmt"
)

//encode a string after rfc1738
func rfc1738_encode(s string) string {
	buf := []byte(s)
	var t []byte

	mustencode := func(b byte) bool {
		if (b < '0' || b > '9') &&
			(b < 'a' || b > 'z') &&
			(b < 'A' || b > 'Z') &&
			b != '.' &&
			b != '-' &&
			b != '_' &&
			b != '~' {
			return true
		}
		return false
	}

	for i := 0; i < len(buf); i++ {
		b := buf[i]
		if mustencode(b) {
			t = append(t, '%')
			s := fmt.Sprintf("%X", b)
			t = append(t, []byte(s)...)
		} else {
			t = append(t, b)
		}
	}

	return string(t)
}

/*func main() {
	s := "\x12\x34\x56\x78\x9a\xbc\xde\xf1\x23\x45\x67\x89\xab\xcd\xef\x12\x34\x56\x78\x9a"
	t := rfc1738_encode(s)

	fmt.Println(t)
}*/
