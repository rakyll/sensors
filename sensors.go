// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sensors provides sensor events from various movement sensors.
package sensors

import (
	"errors"
	"sync"
	"time"
)

// Type represents a sensor type.
type Type int

var (
	Accelerometer = Type(1)
	Gyroscope     = Type(2)
	Magnetometer  = Type(3)
	Altimeter     = Type(4)
)

// Event represents a sensor event.
type Event struct {
	// Sensor is the type of the sensor the event is coming from.
	Sensor Type

	// Timestamp is the time the event happened.
	// Timestamps are not Unix times, they represent a time that is
	// only valid for the device's default accelerometer sensor.
	Timestamp int64

	// Data is the event data.
	//
	// If the event source is Accelerometer,
	//  - Data[0]: acceleration force in x axis in m/s^2
	//  - Data[1]: acceleration force in y axis in m/s^2
	//  - Data[2]: acceleration force in z axis in m/s^2
	//
	// If the event source is Gyroscope,
	//  - Data[0]: rate of rotation around the x axis in rad/s
	//  - Data[1]: rate of rotation around the y axis in rad/s
	//  - Data[2]: rate of rotation around the z axis in rad/s
	//
	// If the event source is Magnetometer,
	//  - Data[0]: force of gravity along the x axis in m/s^2
	//  - Data[1]: force of gravity along the y axis in m/s^2
	//  - Data[2]: force of gravity along the z axis in m/s^2
	//
	Data []float64
}

// Manager multiplexes sensor event data from various sensor sources.
type Manager struct {
	once sync.Once
	m    *manager // platform-specific implementation of the underlying manager
}

func (m *Manager) init() {
	m.m = &manager{}
}

// Enable enables a sensor with the specified delay rate.
// If there are multiple sensors of type t on the device, Enable uses
// the default one.
// If there is no default sensor of type t on the device, an error returned.
// Valid sensor types supported by this package are Accelerometer,
// Gyroscope, Magnetometer and Altimeter.
func (m *Manager) Enable(t Type, delay time.Duration) error {
	m.once.Do(m.init)
	if t < 1 || t > 4 {
		return errors.New("sensors: unknown sensor type")
	}
	return enable(m.m, t, delay)
}

// Disable disables to feed the manager with the specified sensor.
func (m *Manager) Disable(t Type) error {
	m.once.Do(m.init)
	if t < 1 || t > 4 {
		return errors.New("sensors: unknown sensor type")
	}
	return disable(m.m, t)
}

// Read reads a series of events from the manager.
// It may read up to len(e) number of events, but will return
// less events if timeout occurs.
func (m *Manager) Read(e []Event) (n int, err error) {
	m.once.Do(m.init)
	return read(m.m, e)
}

// Close stops the manager and frees the related resources.
func (m *Manager) Close() error {
	m.once.Do(m.init)
	return close(m.m)
}
