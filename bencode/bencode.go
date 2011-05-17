package bencode

import (
	"strconv"
	"os"
)

//a parser struct holding a reference to the token stream and a position pointer
type Parser struct {
	stream []byte
	pos    int
}

//create a new parser for the given token stream
func NewParser(stream []byte) *Parser {
	return &Parser{stream, 0}
}

//wrapper for Parser.ParseOne
//will get one object from the token stream
//use if you know that there's only one object and you won't
//need the input stream anymore as you won't get any position information
func ParseOne(in []byte) (res interface{}, err os.Error) {
	p := NewParser(in)
	return p.ParseOne()
}

//wrapper for Parser.ParseAll
//will return all objects encoded in the token stream
func ParseAll(in []byte) (res []interface{}, err os.Error) {
	p := NewParser(in)
	return p.ParseAll()
}

//return one object from Parser's input stream and advance the pos
func (self *Parser) ParseOne() (res interface{}, err os.Error) {
	return self.nextObject()
}

//return all objects from Parser's input stream and advance to stream end
func (self *Parser) ParseAll() (res []interface{}, err os.Error) {
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
func (self *Parser) nextObject() (res interface{}, err os.Error) {
	switch c := self.stream[self.pos]; {
	case c == 'i':
		return self.nextInteger()
	case c >= '0' && c <= '9':
		return self.nextString()
	case c == 'l':
			return self.nextList()
	}
	return nil, os.NewError("Couldn't parse '" + string(self.stream) + "' ... '" + string(self.stream[self.pos]) + "'")
}

//fetches next integer from stream and advances pos pointer
func (self *Parser) nextInteger() (res int64, err os.Error) {
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
	res, err = strconv.Atoi64(s)
	if err != nil {
		return
	}
	self.pos = idx + 1

	return
}

//fetches next string from stream and advances pos pointer
func (self *Parser) nextString() (res string, err os.Error) {
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

	len_end++	//skip the ':'
	res = string(self.stream[len_end : len_end+l])
	err = nil
	self.pos = len_end+l
	return
}

func (self *Parser) nextList() (res []interface{}, err os.Error) {
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
			self.pos ++ //skip 'e'
			break
		}
	}
	return
}
