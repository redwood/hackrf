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
	"github.com/pkg/errors"
)

func maybeError(code C.int) error {
	if code == HackRFSuccess {
		return nil
	}
	return errors.New(ErrorName(code))
}

// ErrorName function as declared in src/hackrf.h:250
func ErrorName(Errcode C.int) string {
	cstr := C.hackrf_error_name((C.enum_hackrf_error)(Errcode))
	return C.GoString(cstr)
}
