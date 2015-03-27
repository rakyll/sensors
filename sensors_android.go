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
	"fmt"
	"time"
	"unsafe"
)

func init() {
	C.android_initSensors()
}

type manager struct {
	queue *C.ASensorEventQueue
}

func enable(m *manager, t Type, delay time.Duration) error {
	if m.queue == nil {
		q := C.android_createQueue()
		if q == nil {
			return fmt.Errorf("sensors: no default %v sensor", t)
		}
		m.queue = q
	}
	d := delay.Nanoseconds() * 1000
	C.android_enableSensor(m.queue, typeToInt(t), C.int32_t(d))
	return nil
}

func disable(m *manager, t Type) error {
	C.android_disableSensor(m.queue, typeToInt(t))
	return nil
}

func read(m *manager, e []Event) (n int, err error) {
	const size = 5 // number of floats to use per event

	num := len(e)
	dst := make([]float32, size*num)

	n = int(C.android_readQueue(m.queue, C.int(num), (*C.float)(unsafe.Pointer(&dst[0]))))
	for i := 0; i < n/size; i++ {
		ev := Event{}
		// [<type>, <timestamp>, <data>, <data>, <data>]
		ev.Type = intToType(dst[i*size])
		ev.Timestamp = int64(dst[i*size+1])
		ev.Data = []float64{
			float64(dst[i*size+2]),
			float64(dst[i*size+3]),
			float64(dst[i*size+4]),
		}
		e[i] = ev
	}
	return
}

func close(m *manager) error {
	C.android_destroyQueue(m.queue)
	return nil
}

func typeToInt(t Type) C.int {
	switch t {
	case Accelerometer:
		return C.ASENSOR_TYPE_ACCELEROMETER
	case Gyroscope:
		return C.ASENSOR_TYPE_GYROSCOPE
	case Magnetometer:
		return C.ASENSOR_TYPE_MAGNETIC_FIELD
	default:
		return C.int(0)
	}
}

func intToType(t float32) Type {
	switch t {
	case C.ASENSOR_TYPE_ACCELEROMETER:
		return Accelerometer
	case C.ASENSOR_TYPE_GYROSCOPE:
		return Gyroscope
	case C.ASENSOR_TYPE_MAGNETIC_FIELD:
		return Magnetometer
	default:
		return Type(0)
	}
}
