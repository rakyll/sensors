// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,!android darwin,!arm,!arm64

package sensor

import (
	"errors"
	"time"
)

type manager struct{}

func (m *manager) initialize() {}

func (m *manager) enable(s sender, t Type, delay time.Duration) error {
	return errors.New("sensor: no sensors available")
}

func (m *manager) disable(t Type) error {
	return errors.New("sensor: no sensors available")
}

func (m *manager) close() error {
	return nil
}
