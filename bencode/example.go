package main

import (
	"bencode"
	"fmt"
)

func exEncode() {
	fmt.Printf("Encoding example\n================\n")
	enc := bencode.NewEncoder()

	var i int64 = 23
	s := string(bencode.Encode(i))
	fmt.Printf("%d -> %s\n", i, s)
	enc.Encode(i)

	x := "hallo"
	s = string(bencode.Encode(x))
	fmt.Printf("%s -> %s\n", x, s)
	enc.Encode(x)

	var l []interface{}
	l = append(l, int64(44))
	l = append(l, "test")
	s = string(bencode.Encode(l))
	fmt.Printf("list: %#v -> %s\n", l, s)
	enc.Encode(l)

	d := make(map[string]interface{}, 10)
	d["zhort"] = "loli"
	d["ficken"] = 44
	s = string(bencode.Encode(d))
	fmt.Printf("map: %#v -> %s\n", d, s)
	enc.Encode(d)

	fmt.Printf("everything: %s\n", string(enc.Bytes))
}

func main() {
	exEncode()
}
