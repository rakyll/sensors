// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin
// +build arm arm64

package sensor

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreMotion
#import "darwin_armx.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

var (
	aDelay time.Duration
	gDelay time.Duration
	mDelay time.Duration
)

type manager struct {
	m unsafe.Pointer
}

func (m *manager) initialize() {
	m.m = unsafe.Pointer(C.GoIOS_createManager())
}

func (m *manager) enable(t Type, delay time.Duration) error {
	// TODO(jbd): set the interval.
	switch t {
	case Accelerometer:
		C.GoIOS_startAccelerometer(m.m)
		aDelay = delay
	case Gyroscope:
	case Magnetometer:
	default:
		return fmt.Errorf("sensor: unknown sensor type: %v", t)
	}
	return nil
}

func (m *manager) disable(t Type) error {
	switch t {
	case Accelerometer:
		C.GoIOS_stopAccelerometer(m.m)
	case Gyroscope:
	case Magnetometer:
	default:
		return fmt.Errorf("sensor: unknown sensor type: %v", t)
	}
	return nil
}

func (m *manager) read(ch chan interface{}) {
	// TODO(jbd): Disable polling when sensor is disabled.
	go func() {
		ev := make([]C.float, 4)
		var lastTimestamp int64
		for {
			C.GoIOS_readAccelerometer(m.m, (*C.float)(unsafe.Pointer(&ev[0])))
			t := int64(ev[0] * 1000 * 1000)
			if t > lastTimestamp {
				ch <- Event{
					Sensor:    Accelerometer,
					Timestamp: t,
					Data:      []float64{float64(ev[1]), float64(ev[2]), float64(ev[3])},
				}
				lastTimestamp = t
				<-time.Tick(aDelay)
			} else {
				<-time.Tick(aDelay / 2)
			}
		}
	}()
}

func (m *manager) close() error {
	C.GoIOS_destroyManager(m.m)
	return nil
}
