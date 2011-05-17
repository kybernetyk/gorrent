package bencode

import (
	"strconv"
	"os"
)

//a parser struct holding a reference to the token stream and a position pointer
type Decoder struct {
	stream   []byte
	pos      int
	Consumed bool
}

type List []interface{}
type String string
type Integer int64
type Dict map[string]interface{}

//create a new parser for the given token stream
func NewDecoder(stream []byte) *Decoder {
	return &Decoder{stream, 0, false}
}

//return one object from Decoder's input stream and advance the pos
func (self *Decoder) Decode() (res interface{}, err os.Error) {
	return self.nextObject()
}

//return all objects from Decoder's input stream and advance to stream end
func (self *Decoder) DecodeAll() (res []interface{}, err os.Error) {
	for {
		o, e := self.nextObject()
		if e != nil {
			err = e
			return
		}
		res = append(res, o)
		if self.pos >= len(self.stream) {
			break
		}
	}
	return
}

//fetch the next object at position 'pos' in 'stream'
func (self *Decoder) nextObject() (res interface{}, err os.Error) {
	if self.Consumed {
		return nil, os.NewError("This parser's token stream is consumed!")
	}

	switch c := self.stream[self.pos]; {
	case c == 'i':
		res, err = self.nextInteger()
	case c >= '0' && c <= '9':
		res, err = self.nextString()
	case c == 'l':
		res, err = self.nextList()
	default:
		res = nil
		err = os.NewError("Couldn't parse '" + string(self.stream) + "' ... '" + string(self.stream[self.pos]) + "'")
	}
	if self.pos >= len(self.stream) {
		self.Consumed = true
	}
	return
}

//fetches next integer from stream and advances pos pointer
func (self *Decoder) nextInteger() (res Integer, err os.Error) {
	if self.stream[self.pos] != 'i' {
		return 0, os.NewError("No starting 'i' found")
	}

	idx := self.pos + 1
	for {
		if self.stream[idx] == 'e' {
			break
		}
		idx++
		if idx >= len(self.stream) {
			return 0, os.NewError("No ending 'e' found")
		}
	}

	s := string(self.stream[self.pos+1 : idx])
	r, err := strconv.Atoi64(s)
	res = Integer(r)
	if err != nil {
		return
	}
	self.pos = idx + 1

	return
}

//fetches next string from stream and advances pos pointer
func (self *Decoder) nextString() (res String, err os.Error) {
	if self.stream[self.pos] < '0' || self.stream[self.pos] > '9' {
		err = os.NewError("No string length determinator found")
		return
	}

	len_start := self.pos
	len_end := self.pos

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
	self.pos = len_end + l
	return
}

//fetches a list (and its contents) from stream and advances pos
func (self *Decoder) nextList() (res List, err os.Error) {
	if self.stream[self.pos] != 'l' {
		err = os.NewError("This is not a list!")
		return
	}

	self.pos++ //skip 'l'
	for {
		o, e := self.nextObject()
		if e != nil {
			err = e
			return
		}
		res = append(res, o)
		if self.stream[self.pos] == 'e' {
			self.pos++ //skip 'e'
			break
		}
	}
	return
}
