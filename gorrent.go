package main

import (
	"bencode"
	"fmt"
)

/*func ui64() {
	in := "i23e"
	var out int64
	err := bencode.Unmarshal([]byte(in), &out)
	if err != nil {
		fmt.Printf("unmarhsal err: %s\n", err.String())
		return
	}
	fmt.Printf("%s unmarshaled to %d\n", in, out)
}

func uis() {
	in := "4:longtestlol"
	var out string
	err := bencode.Unmarshal([]byte(in), &out)
	if err != nil {
		fmt.Printf("unmarshal err: %s\n", err.String())
		return
	}
	fmt.Printf("%s unmarshaled to '%s'\n", in, out)
}

func uil() {
	in := "l4:testi23e4:leone"
	var out []interface{}
	err := bencode.Unmarshal([]byte(in), &out)
	if err != nil {
		fmt.Printf("unmarshal err: %s\n", err.String())
		return
	}
	fmt.Printf("%s unmarshaled to %#v\n", in, out)
}*/

func main() {
	//	in := "i23ei55e4:test"
	in := "i88eli23ei55e4:testeli33e3:lole"
	//in = "lli4ei5eeli6ei7eee"

	p := bencode.NewParser([]byte(in))

	for !p.Consumed {
		l, err := p.ParseNext()
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
