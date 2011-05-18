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
		return encodeString(in.(string))
	case reflect.Int64:
		return encodeInteger(in.(int64))
	case reflect.Int:
		i := int64(in.(int))
		return encodeInteger(i)
	case reflect.Slice:
		return encodeList(in.([]interface{}))
	case reflect.Map:
		return encodeDict(in.(map[string]interface{}))
	}

	panic("Can't encode this type: " + reflect.TypeOf(in).Name())
}

func encodeString(s string) []byte {
	l := len(s)
	ret := fmt.Sprintf("%d:%s", l, s)
	return []byte(ret)
}

func encodeInteger(i int64) []byte {
	ret := fmt.Sprintf("i%de", i)
	return []byte(ret)
}

func encodeList(list []interface{}) []byte {
	ret := []byte("l")
	for i := 0; i < len(list); i++ {
		o := list[i]
		ret = append(ret, encodeObject(o)...)
	}
	ret = append(ret, 'e')
	return ret
}

func encodeDict(m map[string]interface{}) []byte {
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
