package main

import (
	"C"

	ole "github.com/go-ole/go-ole"
	"github.com/gofrs/uuid"
)
import "fmt"

func main() {
}

//export uuid_v1A
func uuid_v1A() *C.char {
	guid, _ := uuid.NewV1()
	return C.CString(fmt.Sprintf("%v", guid))
}

//export uuid_v1W
func uuid_v1W() *int16 {
	guid, _ := uuid.NewV1()
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}

//export uuid_v4A
func uuid_v4A() *C.char {
	guid, _ := uuid.NewV4()
	return C.CString(fmt.Sprintf("%v", guid))
}

//export uuid_v4W
func uuid_v4W() *int16 {
	guid, _ := uuid.NewV4()
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}

//export uuid_v6A
func uuid_v6A() *C.char {
	guid, _ := uuid.NewV6()
	return C.CString(fmt.Sprintf("%v", guid))
}

//export uuid_v6W
func uuid_v6W() *int16 {
	guid, _ := uuid.NewV6()
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}

//export uuid_v7A
func uuid_v7A() *C.char {
	guid, _ := uuid.NewV7(uuid.MicrosecondPrecision)
	return C.CString(fmt.Sprintf("%v", guid))
}

//export uuid_v7W
func uuid_v7W() *int16 {
	guid, _ := uuid.NewV7(uuid.MicrosecondPrecision)
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}
