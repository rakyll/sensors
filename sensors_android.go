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
	"errors"
	"time"
)
import "unsafe"

func init() {
	C.initSensors()
}

func startAccelerometer(samplesPerSec int) error {
	if ecode := C.startAccelerometer(C.int(samplesPerSec)); ecode == C.ENOSENSOR {
		return errors.New("sensors: no accelerometer sensor on the device")
	}
	return nil
}

func pollAccelerometer(n int) []AccelerometerEvent {
	r := make([]AccelerometerEvent, n)

	var ptr *C.AccelerometerEvent
	ptr = C.pollAccelerometer(C.int(n))

	start := unsafe.Pointer(ptr)
	var item C.AccelerometerEvent

	for i := 0; i < n; i++ {
		current := (*C.AccelerometerEvent)(unsafe.Pointer(uintptr(start) + uintptr(i)*unsafe.Sizeof(item)))
		if current == nil {
			break
		}
		r[i] = AccelerometerEvent{
			DeltaX:    float64(current.x),
			DeltaY:    float64(current.y),
			DeltaZ:    float64(current.z),
			CreatedAt: time.Unix(int64(current.timestamp), 0),
		}
	}

	C.free(start)
	return r
}

func stopAccelerometer() error {
	C.destroyAccelerometer()
	return nil
}
