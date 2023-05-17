package main

import (
	"fmt"
	"unsafe"

	"github.com/valyala/fastjson"
)

//export level_mapper
func level_mapper(tag *uint8, tag_len uint, time_sec uint, time_nsec uint, record *uint8, record_len uint) *uint8 {
	brecord := unsafe.Slice(record, record_len)

	br := string(brecord)
	var p fastjson.Parser
	value, err := p.Parse(br)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	obj, err := value.Object()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Get the level value and convert it to the string
	level := string(obj.Get("level").GetStringBytes())

	// Syslog severity levels https://datatracker.ietf.org/doc/html/rfc3164#autoid-8
	var mappedLevel string
	switch level {
	case "0":
		mappedLevel = "FATAL"
	case "1", "2", "3":
		mappedLevel = "ERROR"
	case "4", "5":
		mappedLevel = "WARN"
	case "6":
		mappedLevel = "INFO"
	case "7":
		mappedLevel = "DEBUG"
	default:
		mappedLevel = "LEVEL_UNSPECIFIED"
	}

	var arena fastjson.Arena
	// Set new key if you want to save original value
	obj.Set("mappedLevel", arena.NewString(mappedLevel))
	// or replace it
	obj.Set("level", arena.NewString(mappedLevel))
	s := obj.String()
	s += string(rune(0)) // Note: explicit null terminator.
	rv := []byte(s)

	return &rv[0]
}

func main() {}
