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

var nextLooperID int

func init() {
	C.android_initSensors()
}

type sensor struct {
	kind     int
	looperId int
	queue    *C.ASensorEventQueue
}

func startAccelerometer(delay time.Duration) (interface{}, error) {
	return startSensor(C.ASENSOR_TYPE_ACCELEROMETER, delay.Nanoseconds()*1000)
}

func closeAccelerometer(s interface{}) error {
	return closeSensor(s.(*sensor))
}

func readAccelerometer(s interface{}, events [][]float64) (n int, err error) {
	return readSensor(s.(*sensor), events)
}

func startSensor(typ int, delay int64) (*sensor, error) {
	id := nextLooperID
	q := C.android_startSensorQueue(C.int(id), C.int(typ), C.int32_t(delay))
	if q == nil {
		return nil, errors.New("sensors: cannot find the default sensor on the device")
	}
	return &sensor{kind: typ, looperId: id, queue: q}, nil
}

func readSensor(s *sensor, events [][]float64) (n int, err error) {
	num := len(events)
	ptr := C.android_readSensorQueue(C.int(s.looperId), s.queue, C.int(num))
	var item C.SensorEvent
	for i := 0; i < num; i++ {
		n = i
		current := (*C.SensorEvent)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i)*unsafe.Sizeof(item)))
		if current == nil {
			break
		}
		events[i] = []float64{
			float64(current.vals[0]),
			float64(current.vals[1]),
			float64(current.vals[2]),
		}
	}
	C.free(unsafe.Pointer(ptr))
	return
}

func closeSensor(s *sensor) error {
	C.android_destroySensorQueue(s.queue)
	return nil
}
