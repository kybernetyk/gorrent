package bencode

import (
	"strconv"
	"os"
)

func Unmarshal(in []byte, out interface{}) os.Error {
	//should we switch based on reflect type of out instead?
	switch c := in[0]; {
	case c == 'i':
		return unmarshalInt64(in, out.(*int64))
	case c >= '0' && c <= '9':
		return unmarshalString(in, out.(*string))
	}

	return os.NewError("Couldn't make sense of '" + string(in) + "' ...") 
}

func unmarshalInt64(in []byte, out *int64) os.Error {
	*out = 0
	if in[0] != 'i' {
		return os.NewError("No starting 'i' found")
	}

	idx := 1
	for {
		if in[idx] == 'e' {
			break
		}
		idx++
		if idx >= len(in) {
			return os.NewError("No ending 'e' found")
		}
	}

	s := string(in[1:idx])
	r, err := strconv.Atoi64(s)
	if err != nil {
		return err
	}
	*out = r
	return nil
}

func unmarshalString(in []byte, out *string) os.Error {
	if in[0] < '0' || in[0] > '9' {
		return os.NewError("No leading length specifier found")
	}

	var len_start int = 0
	var len_end int = 0

	//let's scan the length
	for {
		if in[len_end] == ':' {
			break
		}
		len_end++
		if len_end >= len(in) {
			return os.NewError("No string found.")
		}
	}

	l, err := strconv.Atoi(string(in[len_start:len_end]))
	if err != nil {
		return err
	}
	if l >= len(in[len_end:]) {
		return os.NewError("Length specifier longer than string data")
	}
	//skip the ':'
	len_end++

	s := string(in[len_end:len_end+l])
	*out = s
	return nil
}
