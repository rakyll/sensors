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

type Type int

var (
	Accelerometer = Type(1)
	Gyroscope     = Type(2)
	Magnetometer  = Type(3)
	Altimeter     = Type(4)
)

type Event struct {
	Type      Type
	Timestamp int64
	Data      []float64
}

type Manager struct {
	once sync.Once
	m    *manager // platform-specific implementation of the underlying manager
}

func (m *Manager) init() {
	m.m = &manager{}
}

func (m *Manager) Enable(t Type, delay time.Duration) error {
	m.once.Do(m.init)
	if t < 1 || t > 4 {
		return errors.New("sensors: unknown sensor type")
	}
	return enable(m.m, t, delay)
}

func (m *Manager) Disable(t Type) error {
	m.once.Do(m.init)
	if t < 1 || t > 4 {
		return errors.New("sensors: unknown sensor type")
	}
	return disable(m.m, t)
}

func (m *Manager) Read(e []Event) (n int, err error) {
	m.once.Do(m.init)
	return read(m.m, e)
}

// Close stops the manager and frees the related resources.
func (m *Manager) Close() error {
	m.once.Do(m.init)
	return close(m.m)
}
