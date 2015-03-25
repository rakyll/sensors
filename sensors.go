// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sensors provides sensor events from various movement and position sensors.
package sensors

import "time"

type Type int

var (
	Accelerometer = Type(1)
	Gyroscope     = Type(2)
	Magnometer    = Type(3)
	Altimeter     = Type(4)
)

type Manager struct {
	m interface{} // platform-specific implementation of the underlying manager
}

func (m *Manager) Enable(t Type, delay time.Duration) error {
	panic("not yet implemented")
}

func (m *Manager) Disable(t Type) error {
	panic("not yet implemented")
}

func (a *Manager) Read(e [][]float64) (n int, err error) {
	panic("not yet implemented")
}

// Close stops the manager and frees the related resources.
func (m *Manager) Close() error {
	panic("not yet implemented")
}
