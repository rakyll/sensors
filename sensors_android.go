// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sensors

/*
#cgo LDFLAGS: -llog -landroid

#include <stdlib.h>
#include <android/sensor.h>

#include "sensors_android.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)
import "time"

var (
	aStop chan struct{}
)

func init() {
	C.initSensors()
}

func startAccelerometer(fn func(deltaX, deltaY, deltaZ float64)) {
	aStop = make(chan struct{})
	go func() {
		// TODO(jbd): Need to runtime.LockOSThread?
		runtime.LockOSThread()
		C.startAccelerometer()
		for {
			select {
			case <-aStop:
				return
			default:
				ev := C.pollAccelerometer()
				fn(float64(ev.x), float64(ev.y), float64(ev.z))
				C.free(unsafe.Pointer(ev))
			}
			time.Sleep(time.Microsecond)
		}
	}()
}

func stopAccelerometer() {
	aStop <- struct{}{}
	C.destroyAccelerometer()
}
