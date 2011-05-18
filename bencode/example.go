package main

import (
"bencode"
"fmt"
)

func main() {
	var i int64 = 23
	s := string(bencode.Encode(i))
	fmt.Printf("%d -> %s\n", i, s)

	x := "hallo"
	s = string(bencode.Encode(x))
	fmt.Printf("%s -> %s\n", x, s)

	var l []interface{}
	l = append(l, int64(44))
	l = append(l, "test")
//	fmt.Printf("l %#v", l)
	s = string(bencode.Encode(l))
	fmt.Printf("list: %s\n", s)

	d := make(map[string]interface{}, 10)
	d["zhort"] = "loli"
	d["ficken"] = 44
	s = string(bencode.Encode(d))
	fmt.Printf("map: %s\n", s)

}
