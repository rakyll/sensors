// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sensors

type Accelerometer struct {
	fn func(deltaX, deltaY, deltaZ float64)
}

func NewAccelerometer(distance float64, fn func(deltaX, deltaY, deltaZ float64)) (*Accelerometer, error) {
	panic("not yet implemented")
}

func (a *Accelerometer) Stop() error {
	panic("not yet implemented")
}

type Gyroscope struct {
	fn func(roll, pitch, yaw float64)
}

func NewGyroscope(fn func(roll, pitch, yaw float64)) (*Gyroscope, error) {
	panic("not yet implemented")
}

func (g *Gyroscope) Stop() error {
	panic("not yet implemented")
}

type Magnetometer struct {
	fn func(azimut, pitch, roll float64)
}

func NewMagnetometer(fn func(azimut, pitch, roll float64)) (*Magnetometer, error) {
	panic("not yet implemented")
}

func (m *Magnetometer) Stop() error {
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
	panic("not yet")
}
