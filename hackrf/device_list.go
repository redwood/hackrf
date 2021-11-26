package hackrf

/*
#cgo CFLAGS: -msse4.1 -Wall -fPIC -I. -I.. -I../host/libhackrf/src
#cgo LDFLAGS: -lm -lusb-1.0 ${SRCDIR}/../host/build/libhackrf/src/libhackrf.a
#include <inttypes.h>
#include <complex.h>
#include "host/libhackrf/src/hackrf.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
// import "C"
// import (
// 	"runtime"
// 	"unsafe"
// )

// // DeviceList as declared in src/hackrf.h:167
// type DeviceList struct {
// 	SerialNumbers  [][]byte
// 	USBBoardIds    []USBBoardId
// 	USBDeviceIndex []int32
// 	Devicecount    int32
// 	USBDevices     []unsafe.Pointer
// 	USBDevicecount int32
// 	ptr            *C.hackrf_device_list_t
// }

// // GetDeviceList function as declared in src/hackrf.h:182
// func GetDeviceList() *DeviceList {
// 	clist := C.hackrf_device_list()

// 	// serialNumbers := make([][]byte, len(clist.serial_numbers))
// 	// for i := range serialNumbers {
// 	//     serialNumbers[i] = C.GoBytes(clist.serial_numbers[i], len())
// 	// }

// 	list := &DeviceList{
// 		// SerialNumbers: serialNumbers,
// 		USBDeviceCount: int32(clist.usb_devicecount),
// 		ptr:            clist,
// 	}
// 	runtime.SetFinalizer(list, func() {
// 		C.hackrf_device_list_free(list.ptr)
// 	})
// 	return list
// }

// func (dl *DeviceList) OpenByIndex(i int) (*Device, error) {
// 	C.hackrf_device_list_open()
// }
