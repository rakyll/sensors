// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sensors

import "time"

// TODO(jbd): Find a better name for HandleFunc.

type Location struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	Accuracy  float64
}

type LocationNotifier struct {
	distance float64
	fn       func(loc *Location, at time.Time)
}

func NewLocationNotifier(distance float64) (*LocationNotifier, error) {
	panic("not yet")
}

func (l *LocationNotifier) Stop() error {
	panic("not yet")
}

func (l *LocationNotifier) HandleFunc(fn func(loc *Location, at time.Time)) {
	l.fn = fn
}

func (l *LocationNotifier) LastKnownLocation() (*Location, error) {
	panic("not yet")
}

type Accelerometer struct {
	fn func(deltaX, deltaY, deltaZ float64)
}

func NewAccelerometer() (*Accelerometer, error) {
	panic("not yet implemented")
}

func (a *Accelerometer) HandleFunc(fn func(deltaX, deltaY, deltaZ float64)) {
	a.fn = fn
}

func (a *Accelerometer) Stop() error {
	panic("not yet implemented")
}

type Gyroscope struct {
	fn func(roll, pitch, yaw float64)
}

func NewGyroscope() (*Gyroscope, error) {
	panic("not yet implemented")
}

func (g *Gyroscope) HandleFunc(fn func(roll, pitch, yaw float64)) {
	g.fn = fn
}

func (g *Gyroscope) Stop() error {
	panic("not yet implemented")
}

type Magnetometer struct {
	fn func(azimut, pitch, roll float64)
}

func NewMagnetometer() (*Magnetometer, error) {
	panic("not yet implemented")
}

func (m *Magnetometer) HandleFunc(fn func(azimut, pitch, roll float64)) {
	m.fn = fn
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
