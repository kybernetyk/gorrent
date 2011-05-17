package main

import (
	"bencode"
	"fmt"
	"reflect"
)

func main() {
	//	in := "i23ei55e4:test"
	in := "i88eli23ei55e4:testeli33e3:lole"
	//in = "lli4ei5eeli6ei7eee"
	p := bencode.NewDecoder([]byte(in))

	for !p.Consumed {
		l, err := p.Decode()
		if err != nil {
			fmt.Printf("parser error: %s\n", err.String())
			break
		}
		switch l.(type) {
		case bencode.List:
			x := l.(bencode.List)
			fmt.Printf("list:\n")
			for _, o := range x {
				fmt.Printf("\tobj(%s): %#v\n", reflect.TypeOf(o).Name(), o)
			}
			fmt.Printf("list_end\n")
		default:
			fmt.Printf("obj(%s): %#v\n", reflect.TypeOf(l).Name(), l)
		}
	}
}
