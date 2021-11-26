package hackrf

/*
#cgo CFLAGS: -msse4.1 -Wall -fPIC -I. -I.. -I../host/libhackrf/src
#cgo LDFLAGS: -lm -lusb-1.0 ${SRCDIR}/../host/build/libhackrf/src/libhackrf.a
#include <inttypes.h>
#include <complex.h>
#include "host/libhackrf/src/hackrf.h"
#include <stdlib.h>
#include "cgo_helpers.h"
#include "extras.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/pkg/errors"
)

//export rxCallbackHelper
func rxCallbackHelper(t *C.hackrf_transfer) int32 {
	if t.rx_ctx == nil {
		fmt.Println("ERROR: could not get context")
		return -1
	}
	return callbackHelper(t, t.rx_ctx)
}

//export txCallbackHelper
func txCallbackHelper(t *C.hackrf_transfer) int32 {
	if t.tx_ctx == nil {
		fmt.Println("ERROR: could not get context")
		return -1
	}
	return callbackHelper(t, t.tx_ctx)
}

func callbackHelper(t *C.hackrf_transfer, handle unsafe.Pointer) int32 {
	callback, is := pointerHandles.Get(handle).(TransferCallback)
	if !is {
		callback(nil, errors.New("could not get callback"))
		return -1
	}

	transfer := &Transfer{
		Device:       &Device{ptr: t.device},
		Buffer:       C.GoBytes(unsafe.Pointer(t.buffer), t.buffer_length),
		BufferLength: int32(t.buffer_length),
		ValidLength:  int32(t.valid_length),
		// RxCtx:        unsafe.Pointer(t.rx_ctx),
		// TxCtx:        unsafe.Pointer(t.tx_ctx),
	}
	return callback(transfer, nil)
}
