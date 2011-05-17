package bencode

import (
	"strconv"
	"os"
)

func Unmarshal(in []byte) (res []interface{}, err os.Error) {
	var idx int64 = 0
	for {
		var br int64 = 0
		if idx >= int64(len(in)) {
			break
		}
		c := in[idx]

		//integer
		if c == 'i' {
			var o int64
			br, err = unmarshalInt64(in[idx:], &o)
			if err != nil {
				return
			}
			res = append(res, o)
			idx += br
			continue
		}

		//string
		if c >= '0' && c <= '9' {
			var o string
			br, err = unmarshalString(in[idx:], &o)
			if err != nil {
				return
			}
			res = append(res, o)
			idx += br
			continue
		}

		//list
		if c == 'l' {
			var o []interface{} = make([]interface{}, 10)
			br, err = unmarshalList(in[idx:], &o)
			if err != nil {
				return
			}
			res = append(res, o)
			idx += br
			continue
		}
	}
	return
}

// unmarshals an bencoded int into out (int64) 
// returns number of consumed bytes + err
func unmarshalInt64(in []byte, out *int64) (bytes_read int64, err os.Error) {
	*out = 0
	if in[0] != 'i' {
		return 0, os.NewError("No starting 'i' found")
	}

	idx := 1
	for {
		if in[idx] == 'e' {
			break
		}
		idx++
		if idx >= len(in) {
			return 0, os.NewError("No ending 'e' found")
		}
	}

	s := string(in[1:idx])
	r, err := strconv.Atoi64(s)
	if err != nil {
		return 0, err
	}
	*out = r
	bytes_read = int64(idx + 1)
	err = nil
	return
}

//unmarshal an bencoded bytestring
//returns number of consumed bytes
func unmarshalString(in []byte, out *string) (bytes_read int64, err os.Error) {
	if in[0] < '0' || in[0] > '9' {
		return 0, os.NewError("No leading length specifier found")
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
			return 0, os.NewError("No string found.")
		}
	}

	l, err := strconv.Atoi(string(in[len_start:len_end]))
	if err != nil {
		return 0, err
	}
	if l >= len(in[len_end:]) {
		return 0, os.NewError("Length specifier longer than string data")
	}
	//skip the ':'
	len_end++

	s := string(in[len_end : len_end+l])
	*out = s
	bytes_read = int64(len_end + l)
	err = nil
	return
}

func unmarshalList(in []byte, out *[]interface{}) (bytes_read int64, err os.Error) {
	return
}
