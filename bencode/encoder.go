package bencode

import (
	"fmt"
	"reflect"
	"sort"
)

//Encoder takes care of encoding objects into byte streams.
//The result of the encoding operation is available in Encoder.Bytes.
//Consecutive operations are appended to the byte stream.
//
//Accepts only string, int/int64, []interface{} and map[string]interface{} as input.
type Encoder struct {
	Bytes []byte		//the result byte stream
}

func NewEncoder() *Encoder {
	return &Encoder{}
}

//Encode is a wrapper for Encoder.Encode.
//It returns the bencoded byte stream.
func Encode(in interface{}) []byte {
	enc := NewEncoder()
	enc.Encode(in)
	return enc.Bytes
}

//Encode encodes an object into a bencoded byte stream.
//The result of the operation is accessible through Encoder.Bytes.
//
//Example:
//	enc.Encode(23)
//	enc.Encode("test")
//	enc.Result //contains 'i23e4:test'
func (enc *Encoder) Encode(in interface{}) {
	b := enc.encodeObject(in)
	if len(b) > 0 {
		enc.Bytes = append(enc.Bytes, b...)
	}
}

func (enc *Encoder) encodeObject(in interface{}) []byte {
	switch reflect.TypeOf(in).Kind() {
	case reflect.String:
		return enc.encodeString(in.(string))
	case reflect.Int64:
		return enc.encodeInteger(in.(int64))
	case reflect.Int:
		i := int64(in.(int))
		return enc.encodeInteger(i)
	case reflect.Slice:
		return enc.encodeList(in.([]interface{}))
	case reflect.Map:
		return enc.encodeDict(in.(map[string]interface{}))
	default:
		panic("Can't encode this type: " + reflect.TypeOf(in).Name())
	}
	return nil
}

func (enc *Encoder) encodeString(s string) []byte {
	l := len(s)
	if l <= 0 {
		return nil
	}
	ret := fmt.Sprintf("%d:%s", l, s)
	return []byte(ret)
}

func (enc *Encoder) encodeInteger(i int64) []byte {
	ret := fmt.Sprintf("i%de", i)
	return []byte(ret)
}

func (enc *Encoder) encodeList(list []interface{}) []byte {
	if len(list) <= 0 {
		return nil
	}
	ret := []byte("l")
	for i := 0; i < len(list); i++ {
		o := list[i]
		ret = append(ret, enc.encodeObject(o)...)
	}
	ret = append(ret, 'e')
	return ret
}

func (enc *Encoder) encodeDict(m map[string]interface{}) []byte {
	if len(m) <= 0 {
		return nil
	}
	//sort the map >.<
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.SortStrings(keys)

	ret := []byte("d")
	for _, k := range keys {
		v := m[k]
		ret = append(ret, enc.encodeString(k)...)
		ret = append(ret, enc.encodeObject(v)...)
	}
	ret = append(ret, 'e')
	return ret
}
