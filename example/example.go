// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux

package main

import (
	"time"

	"github.com/rakyll/sensors"
	"golang.org/x/mobile/app"
)

func main() {
	sensor.Enable(sensor.Accelerometer, time.Millisecond)
	sensor.Enable(sensor.Gyroscope, time.Second)

	go func() {
		<-time.Tick(time.Second)
		sensor.Disable(sensor.Accelerometer)
	}()

	app.Main(func(a app.App) {})
}
