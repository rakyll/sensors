// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux

package main

import (
	"log"
	"time"

	"github.com/rakyll/sensors"
	"golang.org/x/mobile/app"
)

func main() {
	sensor.Enable(sensor.Accelerometer, time.Microsecond)
	sensor.Enable(sensor.Gyroscope, time.Microsecond)
	go func() {
		for e := range sensor.Events() {
			log.Printf("%v\n", e)
		}
	}()

	go func() {
		<-time.Tick(time.Second)
		// sensor.Disable(sensor.Gyroscope)
	}()

	app.Main(func(a app.App) {})
}
