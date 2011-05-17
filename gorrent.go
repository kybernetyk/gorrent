package main

import (
	"bencode"
	"fmt"
	"strings"
)

func main() {
	//	in := "i23ei55e4:test"
	in := "i88eli23ei55e4:testeli33e3:lole"
	//in = "lli4ei5eeli6ei7eee"
	r := strings.NewReader(in)
	p := bencode.NewDecoder(r)

	for !p.Consumed {
		l, err := p.Decode()
		if err != nil {
			fmt.Printf("parser error: %s\n", err.String())
			break
		}
		switch l.(type) {
		case bencode.List:
			x := l.(bencode.List)
			for _, o := range x {
				fmt.Printf("sbject: %#v\n", o)
			}
		default:
			fmt.Printf("object: %#v\n", l)
		}
	}
}
