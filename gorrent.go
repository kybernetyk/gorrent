package main

import (
	"bencode"
	"fmt"
	"reflect"
//	"io/ioutil"
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
		case []interface{}:
			x := l.([]interface{})
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
	/*in := "i23ei55e4:test"
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
	dict := r.(bencode.Dict)
	info := dict["info"].(bencode.Dict)
	pieces := info["pieces"].(bencode.String)
	piece := pieces[0:20]
	fmt.Printf("len: %d\n", len(piece))
	fmt.Printf("piece: %v\n", []byte(piece))
	fmt.Printf("enc: %s\n", rfc1738_encode(string(piece)))
	fmt.Printf("%s\n", rfc1738_encode("\x12\x34\x56\x78\x9a\xbc\xde\xf1\x23\x45\x67\x89\xab\xcd\xef\x12\x34\x56\x78\x9a"))
	fmt.Printf("%s\n", "%124Vx%9A%BC%DE%F1%23Eg%89%AB%CD%EF%124Vx%9A")
*/
	torrent := &MetaInfo{}
	err := torrent.ReadFromFile("test.torrent")
	if err != nil {
		fmt.Println(err.String())
		return
	}

	//fmt.Printf("%#v\n", torrent.parsed)
	b := torrent.InfoHash()
	s := rfc1738_encode(string(b))
	fmt.Printf("%s\n", s)

}

