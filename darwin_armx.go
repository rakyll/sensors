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
	"sync"
	"time"
	"unsafe"
)

var (
	mu    sync.Mutex
	doneA chan struct{}
)

type manager struct {
	m unsafe.Pointer
}

func (m *manager) initialize() {
	m.m = unsafe.Pointer(C.GoIOS_createManager())
}

func (m *manager) enable(s sender, t Type, delay time.Duration) error {
	// TODO(jbd): If delay is smaller than 10 milliseconds, set it to
	// 10 milliseconds. It is highest frequency iOS SDK suppports and
	// we don't want to have time.Tick durations smaller than this value.
	mu.Lock()
	mu.Unlock()

	switch t {
	case Accelerometer:
		if doneA != nil {
			return fmt.Errorf("sensor: cannot enable; %v sensor is already enabled", t)
		}
		// TODO(jbd): Check if acceloremeter is available.
		C.GoIOS_startAccelerometer(m.m)
		m.startAccelometer(s, delay)
		doneA = make(chan struct{})
	case Gyroscope:
	case Magnetometer:
	default:
		return fmt.Errorf("sensor: unknown sensor type: %v", t)
	}
	return nil
}

func (m *manager) disable(t Type) error {
	mu.Lock()
	mu.Unlock()

	switch t {
	case Accelerometer:
		if doneA == nil {
			return fmt.Errorf("sensor: cannot disable; %v sensor is not enabled", t)
		}
		doneA <- struct{}{}
		C.GoIOS_stopAccelerometer(m.m)
		doneA = nil
	case Gyroscope:
	case Magnetometer:
	default:
		return fmt.Errorf("sensor: unknown sensor type: %v", t)
	}
	return nil
}

func (m *manager) startAccelometer(app sender, d time.Duration) {
	go func() {
		ev := make([]C.float, 4)
		var lastTimestamp int64
		for {
			select {
			case <-doneA:
				return
			default:
				C.GoIOS_readAccelerometer(m.m, (*C.float)(unsafe.Pointer(&ev[0])))
				t := int64(ev[0] * 1000 * 1000)
				if t > lastTimestamp {
					// TODO(jbd): Do we need to convert the values to another unit?
					// How does iOS units compate to the Android units.
					app.Send(Event{
						Sensor:    Accelerometer,
						Timestamp: t,
						Data:      []float64{float64(ev[1]), float64(ev[2]), float64(ev[3])},
					})
					lastTimestamp = t
					<-time.Tick(d)
				} else {
					<-time.Tick(d / 2)
				}
			}
		}
	}()
}

// TODO(jbd): Remove close?
func (m *manager) close() error {
	C.GoIOS_destroyManager(m.m)
	return nil
}
