// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sensors provides sensor events from various movement and position sensors.
package sensors

import (
	"errors"
	"sync"
	"time"
)

var (
	muAStarted sync.Mutex
	aStarted   = false
)

// AccelerometerEvent represents data returned from an accelerometer sensor.
type AccelerometerEvent struct {
	// DeltaX is the delta movement in the x-axis.
	DeltaX float64

	// DeltaY is the delta movement in the y-axis.
	DeltaY float64

	// DeltaZ is the delta movement in the z-axis.
	DeltaZ float64

	// Timestamp is the time the current event has happened. Timestamps are not
	// Unix times, they represent a time that's only valid for the device's
	// default accelerometer sensor.
	Timestamp int64
}

type GyroscopeEvent struct {
	Roll      float64
	Pitch     float64
	Yaw       float64
	Timestamp int64
}

type MagnetometerEvent struct {
	Azimut    float64
	Pitch     float64
	Roll      float64
	Timestamp int64
}

// StartAccelerometer starts the accelerometer.
// Once the accelerometer is no longer in use, it should be stopped
// by calling StopAccelerometer.
// Delay determines the wait-time to read the next sample from the sensor.
// Its lower limit is bound by the sensor's output bandwidth.
func StartAccelerometer(delay time.Duration) error {
	muAStarted.Lock()
	defer muAStarted.Unlock()
	if aStarted {
		return errors.New("sensors: accelerometer already started")
	}
	if err := startAccelerometer(delay.Nanoseconds() * 1000); err != nil {
		return err
	}
	aStarted = true
	return nil
}

// PollAccelerometer polls n new events from the accelerometer event queue.
// It will block until n events are available to the sensor event queue.
// A call to StartAccelerometer is mandatory to start the accelerometer
// sensor and initialize its event queue.
// You have to call this function from the same OS thread that the
// accelerometer has been started. Use runtime.LockOSThread to lock the
// current goroutine to a particular OS thread.
func PollAccelerometer(n int) []AccelerometerEvent {
	return pollAccelerometer(n)
}

// StopAccelerometer stops the accelerometer and frees the related resources.
func StopAccelerometer() error {
	muAStarted.Lock()
	defer muAStarted.Unlock()
	if !aStarted {
		return errors.New("sensors: accelerometer hasn't been started")
	}
	if err := stopAccelerometer(); err != nil {
		return err
	}
	aStarted = false
	return nil
}

func StartGyroscope(delay time.Duration) error {
	panic("not yet implemented")
}

func PollGyroscope(n int) (roll, pitch, yaw float64) {
	panic("not yet implemented")
}

func StopGyroscope() error {
	panic("not yet implemented")
}

func StartMagnetometer(delay time.Duration) error {
	panic("not yet implemented")
}

func PollMagnetometer(azimut, pitch, roll float64) {
	panic("not yet implemented")
}

func StopMagnetometer() error {
	panic("not yet implemented")
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
