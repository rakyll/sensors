// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sensors

import (
	"errors"
	"sync"
)

var (
	muAStarted sync.Mutex
	aStarted   = false
)

// StartAccelerometer starts the accelerometer.
// Once the accelerometer is no longer in use, it should be stopped
// by calling StopAccelerometer.
func StartAccelerometer(samplesPerSec int) error {
	muAStarted.Lock()
	defer muAStarted.Unlock()
	if aStarted {
		return errors.New("sensors: accelerometer already started")
	}
	if err := startAccelerometer(samplesPerSec); err != nil {
		return err
	}
	aStarted = true
	return nil
}

// PollAccelerometer polls a new event for the accelerometer event queue.
// It blocks until an event is available.
// A call to StartAccelerometer is mandatory to start the accelerometer
// sensor and initialize its event queue.
// You have to call PollAccelerometer from the same OS thread that the
// accelerometer is started on.
func PollAccelerometer() (deltaX, deltaY, deltaZ float64) {
	return pollAccelerometer()
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

func StartGyroscope(fn func(roll, pitch, yaw float64)) error {
	panic("not yet implemented")
}

func StopGyroscope() {
	panic("not yet implemented")
}

func StartMagnetometer(fn func(azimut, pitch, roll float64)) error {
	panic("not yet implemented")
}

func StopMagnetometer() {
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
