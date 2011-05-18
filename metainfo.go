package main

import (
	"bencode"
	"os"
	"io/ioutil"
	"bytes"
	"fmt"
)
//metainfo file (.torrent file) handling

type MetaInfo struct {
	raw    []byte
	parsed bencode.Dict
}


func (mi *MetaInfo) ReadFromFile(filename string) os.Error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	mi.raw = b

	dec := bencode.NewDecoder(b)
	o, err := dec.Decode()
	if err != nil {
		return os.NewError("Couldn't parse torrent: " + err.String())
	}

	mi.parsed = o.(bencode.Dict)
	return nil
}

//god this will be ugly
func (mi *MetaInfo) infoHash() []byte {
	
	return nil
}
