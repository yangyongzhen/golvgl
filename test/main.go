package main

import (
	"fmt"
	"unsafe"
)

// struct x {
//  int y, z;
// };
//
// int sum(struct x a) {
//  return a.y + a.z;
// }
//
import "C"

type X struct{ Y, Z int32 }

type Lobj C.struct_x

func main() {
	a := &X{0, 1}
	fmt.Println(a, "->", C.sum(*((*C.struct_x)(unsafe.Pointer(a)))))

	var b Lobj
	var c Lobj
	//b.Y = 2
	fmt.Println(b, "->", C.sum(*((*C.struct_x)(unsafe.Pointer(&b)))))

	b.test(&c)
}

func (obj *Lobj) test(b *Lobj) {
	fmt.Println("->", C.sum(*((*C.struct_x)(unsafe.Pointer(b)))))
}
