// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sensors provides sensor events from various movement and position sensors.
package sensors

import "time"

// Accelerometer represents an accelerometer sensor.
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
// It will block until len(e) events are retrieved.
// e is a series of 3-vectors. Each vector contains the acceleration
// force in m/s2 that is applied to the device in x, y and z axes.
//
// 	e[i][0]: acceleration force in x-axis
// 	e[i][1]: acceleration force in y-axis
// 	e[i][2]: acceleration force in z-axis
//
// A call to StartAccelerometer is mandatory to start the accelerometer
// sensor and initialize its event queue.
//
// You have to call this function from the same OS thread that the
// accelerometer has been started. Use runtime.LockOSThread to lock the
// current goroutine to a particular OS thread.
func (a *Accelerometer) Read(e [][]float64) (n int, err error) {
	return readAccelerometer(a.s, e)
}

// Close stops the accelerometer and frees the related resources.
func (a *Accelerometer) Close() error {
	return closeAccelerometer(a.s)
}
