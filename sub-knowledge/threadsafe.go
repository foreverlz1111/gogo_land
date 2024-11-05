package main

import (
	"log"
	"sync"
)

// 线程安全的map

type StringMap struct {
	m sync.Map
}

func (s *StringMap) Store(key, value string) {
	s.m.Store(key, value)
}
func (s *StringMap) Load(key string) (value string, ok bool) {
	v, ok := s.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}
func main() {
	stringmap := new(StringMap)
	log.Println(stringmap)
}
