package main

//#cgo CFLAGS: -I${SRCDIR}/callClib
//#cgo LDFlAGS: ./callC.a
//#include <stdio.h>
//#include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.cHello()

	fmt.Println("Goint to call another C function!")
	myMessage := C.String("This is Mihalis!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")
}