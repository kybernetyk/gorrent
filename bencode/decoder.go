/*
	Package bencode implements reading and writing* of 'bencoded'
	object streams used by the Bittorent protocol.

	* = not implemented yet
*/
package bencode

import (
	"strconv"
	"os"
)

//A Decoder reads and decodes bencoded objects from an input stream.
//It returns objects that are either an "Integer", "String", "List" or "Dict".
//
//Example usage:
//	d := bencode.NewDecoder([]byte("i23e4:testi123e"))
//	for !p.Consumed {
//		o, _ := p.Decode()
//		fmt.Printf("obj(%s): %#v\n", reflect.TypeOf(o).Name, o)
//	}
type Decoder struct {
	stream   []byte
	Pos      int
	Consumed bool //true if we have consumed all tokens
}

type List []interface{}
type String string
type Integer int64
type Dict map[string]interface{}


//NewDecoder creates a new decoder for the given token stream
func NewDecoder(b []byte) *Decoder {
	return &Decoder{b, 0, false}
}

//Decode reads one object from the input stream
func (self *Decoder) Decode() (res interface{}, err os.Error) {
	return self.nextObject()
}

//DecodeAll reads all objects from the input stream
func (self *Decoder) DecodeAll() (res []interface{}, err os.Error) {
	for {
		o, e := self.nextObject()
		if e != nil {
			err = e
			return
		}
		res = append(res, o)
		if self.Pos >= len(self.stream) {
			break
		}
	}
	return
}

//fetch the next object at Position 'pos' in 'stream'
func (self *Decoder) nextObject() (res interface{}, err os.Error) {
	if self.Consumed {
		return nil, os.NewError("This parser's token stream is consumed!")
	}

	switch c := self.stream[self.Pos]; {
	case c == 'i':
		res, err = self.nextInteger()
	case c >= '0' && c <= '9':
		res, err = self.nextString()
	case c == 'l':
		res, err = self.nextList()
	case c == 'd':
		res, err = self.nextDict()
	default:
		res = nil
		err = os.NewError("Couldn't parse '" + string(self.stream) + "' ... '" + string(self.stream[self.Pos]) + "'")
	}
	if self.Pos >= len(self.stream) {
		self.Consumed = true
	}
	return
}

//fetches next integer from stream and advances Pos pointer
func (self *Decoder) nextInteger() (res Integer, err os.Error) {
	if self.stream[self.Pos] != 'i' {
		return 0, os.NewError("No starting 'i' found")
	}
	validstart := false //flag to check for leading 0's
	idx := self.Pos + 1
	for {
		if self.stream[idx] == 'e' {
			break
		}

		if self.stream[idx] == '0' && !validstart {
			err = os.NewError("Leading Zeros are not allowed in bencoded integers!")
			return
		}

		//check for bytes != '-' and '0'..'9'
		if (self.stream[idx] < '0' || self.stream[idx] > '9') && self.stream[idx] != '-' {
			err = os.NewError("Invalid byte '" + string(self.stream[idx]) + "' in encoded integer.")
			return
		} else {
			validstart = true
		}

		idx++
		if idx >= len(self.stream) {
			return 0, os.NewError("No ending 'e' found")
		}
	}

	s := string(self.stream[self.Pos+1 : idx])
	r, err := strconv.Atoi64(s)
	res = Integer(r)
	if err != nil {
		return
	}
	self.Pos = idx + 1

	return
}

//fetches next string from stream and advances Pos pointer
func (self *Decoder) nextString() (res String, err os.Error) {
	if self.stream[self.Pos] < '0' || self.stream[self.Pos] > '9' {
		err = os.NewError("No string length determinator found")
		return
	}

	len_start := self.Pos
	len_end := self.Pos

	//scan length
	for {
		if self.stream[len_end] == ':' {
			break
		}
		len_end++
		if len_end >= len(self.stream) {
			err = os.NewError("No string found ...")
			return
		}
	}

	l, e := strconv.Atoi(string(self.stream[len_start:len_end]))
	if e != nil {
		err = os.NewError("Couldn't parse string length specifier: " + e.String())
		return
	}
	if l >= len(self.stream[len_end:]) {
		err = os.NewError("Specified length longer than data buffer ...")
		return
	}

	len_end++ //skip the ':'
	res = String(self.stream[len_end : len_end+l])
	err = nil
	self.Pos = len_end + l
	return
}

//fetches a list (and its contents) from stream and advances Pos
func (self *Decoder) nextList() (res List, err os.Error) {
	if self.stream[self.Pos] != 'l' {
		err = os.NewError("This is not a list!")
		return
	}

	self.Pos++ //skip 'l'
	for {
		o, e := self.nextObject()
		if e != nil {
			err = e
			return
		}
		res = append(res, o)
		if self.stream[self.Pos] == 'e' {
			self.Pos++ //skip 'e'
			break
		}
	}
	return
}

//fetches a dict
//bencoded dicts must have their keys sorted lexically. but I guess
//we can ignore that and work with unsorted maps. (wtf?! sorted maps ...)
func (self *Decoder) nextDict() (res Dict, err os.Error) {
	if self.stream[self.Pos] != 'd' {
		err = os.NewError("This is not a dict!")
		return
	}
	res = make(Dict)
	self.Pos++ //skip 'd'
	for {
		key, e := self.nextString()
		if e != nil {
			err = e
			return
		}
		val, e := self.nextObject()
		if e != nil {
			err = e
			return
		}
		//fmt.Printf("key: %s\nval: %#v\n", key, val)
		res[string(key)] = val
		if self.stream[self.Pos] == 'e' {
			self.Pos++ //skip 'e'
			break
		}
	}
	return
}
