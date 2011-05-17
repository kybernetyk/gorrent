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
	in := "i23ei55e4:test"
	l, err := bencode.Unmarshal([]byte(in))
	if err != nil {
		fmt.Println("unmarshal failed: %s\n", err.String())
		return
	}
	fmt.Printf("unmarshalled: %#v\n", l)
}

