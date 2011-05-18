package main

import (
	"bencode"
	"fmt"
	"reflect"
	"io/ioutil"
)

func test0r(in string) {
	fmt.Printf("testoring: '%s' ...\n", in)
	p := bencode.NewDecoder([]byte(in))

	for !p.Consumed {
		l, err := p.Decode()
		if err != nil {
			fmt.Printf("\tparser error: %s\n", err.String())
			break
		}
		switch l.(type) {
		case bencode.List:
			x := l.(bencode.List)
			fmt.Printf("\tlist:\n")
			for _, o := range x {
				fmt.Printf("\t\tobj(%s): %#v\n", reflect.TypeOf(o).Name(), o)
			}
			fmt.Printf("\tlist_end\n")
		default:
			fmt.Printf("\tobj(%s): %#v\n", reflect.TypeOf(l).Name(), l)
		}
	}
}


func main() {
	in := "i23ei55e4:test"
	test0r(in)
	in = "i00123e"
	test0r(in)
	in = "i-45e"
	test0r(in)
	in = "ix23e"
	test0r(in)
	in = "i88eli23ei55e4:testeli33e3:lole"
	test0r(in)
	in = "lli4ei5eeli6ei7eee"
	test0r(in)
	in = "d3:cow3:moo4:spam4:eggse"
	test0r(in)
	in = "d4:spaml1:a1:bee1:xd4:fick1:oe"
	test0r(in)

	//test reading a torrent
	fmt.Printf("Parsing 'test.torrent' ...\n")
	b, err := ioutil.ReadFile("test.torrent")
	if err != nil {
		panic("couldn't open test.torrent")
	}
	p := bencode.NewDecoder(b)
	r, err := p.Decode()
	if err != nil {
		fmt.Printf("Couldn't parse torrent: %s\n", err.String())
		return
	}
	fmt.Printf("%#v\n", r)
}
