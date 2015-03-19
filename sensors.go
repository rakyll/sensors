// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sensors provides sensor events from various movement and position sensors.
package sensors

import "time"

type Accelerometer struct {
	s interface{}
}

// StartAccelerometer starts a new accelerometer.
// Once the accelerometer is no longer in use, it should be stopped
// and its resources should be released by calling Close.
// Delay determines the wait-time to read the next sample from the sensor.
// Its lower limit is bound by the sensor's output bandwidth.
func StartAccelerometer(delay time.Duration) (*Accelerometer, error) {
	s, err := startAccelerometer(delay.Nanoseconds() * 1000)
	if err != nil {
		return nil, err
	}
	return &Accelerometer{s: s}, nil
}

// Read reads new events from the accelerometer event queue.
// It will block until len(events) number of events are available to
// the sensor event queue.
//
// A call to StartAccelerometer is mandatory to start the accelerometer
// sensor and initialize its event queue.
//
// You have to call this function from the same OS thread that the
// accelerometer has been started. Use runtime.LockOSThread to lock the
// current goroutine to a particular OS thread.
func (a *Accelerometer) Read(events [][]float64) (n int, err error) {
	return readAccelerometer(a.s, events)
}

// Close stops the accelerometer and frees the related resources.
func (a *Accelerometer) Close() error {
	return closeAccelerometer(a.s)
}

// Type of the network that is currently in use.
const (
	TypeWiFi = iota
	TypeMobile
	TypeOther
)

// Connectivity status.
const (
	StatusUnknown = iota
	StatusConnecting
	StatusConnected
	StatusDisconnecting
)

// Connectivity returns the type and the status of the network that is
// currently in use.
func Connectivity() (networkType int, status int) {
	panic("not yet implemented")
}
