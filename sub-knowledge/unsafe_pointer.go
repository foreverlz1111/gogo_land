package main

import (
	"log"
	"unsafe"
)

type Person struct {
	age  int8
	name []string
}

func main() {
	person := new(Person)
	log.Println(person)
	// unsafe.Pointer： GO的指针地址不参与运算
	age := unsafe.Pointer(uintptr(unsafe.Pointer(person)) + unsafe.Offsetof(person.age))
	name := unsafe.Pointer(uintptr(unsafe.Pointer(person)) + unsafe.Offsetof(person.name))
	*((*int)(age)) = 10
	*((*[]string)(name)) = []string{"a", "b"}

	log.Println(person)
}
