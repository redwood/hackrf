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

// Init function as declared in src/hackrf.h:176
func Init() error {
	err := C.hackrf_init()
	return maybeError(err)
}

// Exit function as declared in src/hackrf.h:177
func Exit() error {
	err := C.hackrf_exit()
	return maybeError(err)
}

type Device struct {
	ptr *C.hackrf_device
}

// Open function as declared in src/hackrf.h:186
func NewDevice() (*Device, error) {
	// device := C.make_hackrf_device()
	var ptr *C.hackrf_device
	if err := C.hackrf_open(&ptr); err != 0 {
		return nil, maybeError(err)
	}
	return &Device{ptr: ptr}, nil
}

// Close function as declared in src/hackrf.h:188
func (d *Device) Close() error {
	err := C.hackrf_close(d.ptr)
	return maybeError(err)
}

// Transfer as declared in src/hackrf.h:140
type Transfer struct {
	Device       *Device
	Buffer       []byte
	BufferLength int32
	ValidLength  int32
	// RxCtx          unsafe.Pointer
	// TxCtx          unsafe.Pointer
}

// TransferCallback type as declared in src/hackrf.h:169
type TransferCallback func(t *Transfer, err error) int32

// StartRx function as declared in src/hackrf.h:190
func (d *Device) StartRx(callback TransferCallback) error {
	handle := pointerHandles.Track(callback)
	err := C.hackrf_start_rx(d.ptr, (*[0]byte)(C.rx_callback_helper), handle)
	return maybeError(err)
}

// StopRx function as declared in src/hackrf.h:191
func (d *Device) StopRx() error {
	err := C.hackrf_stop_rx(d.ptr)
	return maybeError(err)
}

// StartTx function as declared in src/hackrf.h:193
func (d *Device) StartTx(callback TransferCallback) error {
	handle := pointerHandles.Track(callback)
	err := C.hackrf_start_tx(d.ptr, (*[0]byte)(C.tx_callback_helper), handle)
	return maybeError(err)
}

// StopTx function as declared in src/hackrf.h:194
func (d *Device) StopTx() error {
	err := C.hackrf_stop_tx(d.ptr)
	return maybeError(err)
}

// IsStreaming function as declared in src/hackrf.h:197
func (d *Device) IsStreaming() bool {
	is := C.hackrf_is_streaming(d.ptr)
	return is > 0
}

// SetFreq function as declared in src/hackrf.h:224
func (d *Device) SetFreq(freqHz uint) error {
	err := C.hackrf_set_freq(d.ptr, (C.uint64_t)(freqHz))
	return maybeError(err)
}

// SetSampleRate function as declared in src/hackrf.h:231
func (d *Device) SetSampleRate(freqHz float64) error {
	err := C.hackrf_set_sample_rate(d.ptr, (C.double)(freqHz))
	return maybeError(err)
}

// SetAmpEnable function as declared in src/hackrf.h:234
func (d *Device) SetAmpEnable(enabled bool) error {
	var e C.uint8_t
	if enabled {
		e = 1
	}
	err := C.hackrf_set_amp_enable(d.ptr, e)
	return maybeError(err)
}

// SetLnaGain function as declared in src/hackrf.h:239
func (d *Device) SetLnaGain(gain uint32) error {
	err := C.hackrf_set_lna_gain(d.ptr, C.uint32_t(gain))
	return maybeError(err)
}

// SetVgaGain function as declared in src/hackrf.h:242
func (d *Device) SetVgaGain(gain uint32) error {
	err := C.hackrf_set_vga_gain(d.ptr, C.uint32_t(gain))
	return maybeError(err)
}

// SetTxvgaGain function as declared in src/hackrf.h:245
func (d *Device) SetTxvgaGain(gain uint32) error {
	err := C.hackrf_set_txvga_gain(d.ptr, C.uint32_t(gain))
	return maybeError(err)
}

// SetAntennaEnable function as declared in src/hackrf.h:248
func (d *Device) SetAntennaEnable(enabled bool) error {
	var e C.uint8_t
	if enabled {
		e = 1
	}
	err := C.hackrf_set_antenna_enable(d.ptr, e)
	return maybeError(err)
}
