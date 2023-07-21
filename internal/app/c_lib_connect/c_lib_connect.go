package c_lib_connect

// #include <stdio.h>
// #include <stdlib.h>
// #include <example.h>
// static void myPrint1(char* s) {
//   printf("%s\n", s);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func Main() {
	//Пример для подключения из example.h
	cs := C.CString("Hello from stdio")
	C.myPrint(cs)
	C.free(unsafe.Pointer(cs))
	//Пример для подключения из комментария.
	cs1 := C.CString("Hello from stdio")
	C.myPrint1(cs1)
	C.free(unsafe.Pointer(cs1))
	fmt.Printf("Done.")
}

/*
set PATH=%PATH%;D:\msys64\ucrt64\bin\
*/
