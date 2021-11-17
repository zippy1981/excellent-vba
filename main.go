package main

import (
	"C"
	"fmt"
	"log"
	"runtime"
	"os"
	//"syscall"

	ole "github.com/go-ole/go-ole"
	"github.com/gofrs/uuid"
)

const PROCMON_DEBUGGER_HANDLER = "\\\\.\\Global\\ProcmonDebugLogger"
const FILE_DEVICE_PROCMON_LOG uint32 = 0x9535
const IOCTL_EXTERNAL_LOG_DEBUGOUT int32 = -1791655420



func main() {
	logger := get_logger()
	k := 1
	// Just to prove we are not loeaking memory.
	for ; k <= 1000000000; k++ {
		guid := UUIDv4W()
		if (k % 100000 == 0) {
			PrintMemUsage(logger)
			logger.Printf("Iteration %d.\n", k)
		}
		err := ole.SysFreeString(guid)
		
		if(err != nil) {
			logger.Printf("err: %v\n", err)
		}
	}
}

func get_logger() log.Logger {
	//logHandle, procmonErr := os.Create(PROCMON_DEBUGGER_HANDLER)
	//if (procmonErr == nil) {
		logHandle := os.Stdout
//	} else {
//		msg := "Hrmhrm"
//		bstrMsg := ole.SysAllocString(msg)		
//		
//		_ = syscall.DeviceIoControl(syscall.Handle(logHandle.Fd()), IOCTL_EXTERNAL_LOG_DEBUGOUT, *byte(bstrMsg), uint32(len(msg) * 2), nil, nil, nil, nil)
//		ole.SysFreeString(bstrMsg)
//	}
	return *log.New(logHandle, "Go-Ole-Guid-Profiling ", log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage(logger log.Logger) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	logger.Printf("Alloc = %v MiB\tSys = %v MiB\tNumGC = %v", bToMb(m.Alloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) uint64 {
return b / 1024 / 1024
}

//export UUIDv1A
func UUIDv1A() *C.char {
	guid, _ := uuid.NewV1()
	return C.CString(fmt.Sprintf("%v", guid))
}

//export UUIDv1W
func UUIDv1W() *int16 {
	guid, _ := uuid.NewV1()
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}

//export UUIDv4A
func UUIDv4A() *C.char {
	guid, _ := uuid.NewV4()
	return C.CString(fmt.Sprintf("%v", guid))
}

//export UUIDv4W
func UUIDv4W() *int16 {
	guid, _ := uuid.NewV4()
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}

//export UUIDv6A
func UUIDv6A() *C.char {
	guid, _ := uuid.NewV6()
	return C.CString(fmt.Sprintf("%v", guid))
}

//export UUIDv6W
func UUIDv6W() *int16 {
	guid, _ := uuid.NewV6()
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}

//export UUIDv7A
func UUIDv7A() *C.char {
	guid, _ := uuid.NewV7(uuid.MicrosecondPrecision)
	return C.CString(fmt.Sprintf("%v", guid))
}

//export UUIDv7W
func UUIDv7W() *int16 {
	guid, _ := uuid.NewV7(uuid.MicrosecondPrecision)
	return ole.SysAllocString(fmt.Sprintf("%v", guid))
}
