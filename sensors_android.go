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
	C.android_enableSensor(m.queue, intToSensorType(t), C.int32_t(d))
	return nil
}

func intToSensorType(t Type) C.int {
	switch t {
	case Accelerometer:
		return C.ASENSOR_TYPE_ACCELEROMETER
	case Gyroscope:
		return C.ASENSOR_TYPE_GYROSCOPE
	case Magnometer:
		return C.ASENSOR_TYPE_MAGNETIC_FIELD
	}
	return C.int(0)
}

func disable(m *manager, t Type) error {
	C.android_disableSensor(m.queue, intToSensorType(t))
	return nil
}

func read(m *manager, e [][]float64) (n int, err error) {
	num := len(e)
	ptr := C.android_readQueue(m.queue, C.int(num))
	var item C.float
	for i := 0; i < num; i++ {
		n = i
		current := (*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i)*4*unsafe.Sizeof(item)))
		if current == nil {
			break
		}
		e[i] = make([]float64, 4)
		for j := 0; j < 4; j++ {
			c := (C.float)(uintptr(unsafe.Pointer(current)) + uintptr(j)*unsafe.Sizeof(item))
			e[i][j] = float64(c)
		}
	}
	C.free(unsafe.Pointer(ptr))
	return
}

func close(m *manager) error {
	C.android_destroyQueue(m.q)
	return nil
}
