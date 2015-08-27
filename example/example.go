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
	app.Main(func(a app.App) {
		sensor.Enable(a, sensor.Accelerometer, time.Millisecond)
		sensor.Enable(a, sensor.Gyroscope, time.Second)

		go func() {
			<-time.Tick(time.Second)
			sensor.Disable(sensor.Accelerometer)
		}()
	})
}