package main

import "testing"

// 接口的实现为指针
func Convert(value int) []interface{} {
	var slice = make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		slice[i] = value
	}
	return slice
}
func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Convert(20)
		_ = result
	}
}
func TestConvert(t *testing.T) {
	result := Convert(20)
	if len(result) != 100 {
		t.Errorf("expected length 100, got %d", len(result))
	}
}
