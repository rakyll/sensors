// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sensors

// StartAccelerometer starts the accelerometer and notifies fn with
// the positional changes in x, y and z axes.
// Once the accelerometer is no longer in use, it should be stopped
// by calling StopAccelerometer.
func StartAccelerometer(fn func(deltaX, deltaY, deltaZ float64)) error {
	startAccelerometer(fn)
	// TODO(jbd): Return error if no default accelerometer is found.
	return nil
}

// StopAccelerometer stops the accelerometer and frees the related resources.
func StopAccelerometer() {
	stopAccelerometer()
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
