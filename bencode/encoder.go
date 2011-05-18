package bencode

import (
	"fmt"
	"reflect"
	"sort"
)


func Encode(in interface{}) []byte {
	return encodeObject(in)
}

func encodeObject(in interface{}) []byte {
	switch reflect.TypeOf(in).Kind() {
	case reflect.String:
		return encodeString(in)
	case reflect.Int64:
		return encodeInteger(in)
	case reflect.Slice:
		return encodeList(in)
	case reflect.Map:
		return encodeDict(in)
	}

	panic("WTF?")
}

func encodeString(in interface{}) []byte {
	o := in.(string)
	s := string(o)
	l := len(s)

	ret := fmt.Sprintf("%d:%s", l, s)
	return []byte(ret)
}

func encodeInteger(in interface{}) []byte {
	o := in.(int64)
	ret := fmt.Sprintf("i%de", o)
	return []byte(ret)
}

func encodeList(in interface{}) []byte {
	list := in.([]interface{})
	ret := []byte("l")
	for i := 0; i < len(list); i++ {
		o := list[i]
		ret = append(ret, encodeObject(o)...)
	}
	ret = append(ret, 'e')
	return ret
}

func encodeDict(in interface{}) []byte {
	m := in.(map[string]interface{})

	//sort the map >.<
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.SortStrings(keys)

	ret := []byte("d")
	for _, k := range keys {
		v := m[k]
		ret = append(ret, encodeString(k)...)
		ret = append(ret, encodeObject(v)...)
	}
	ret = append(ret, 'e')
	return ret
}
