// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sensors

/*
#cgo LDFLAGS: -llog -landroid

#include <stdlib.h>
#include <android/sensor.h>

#include "sensors_android.h"
*/
import "C"

func init() {
	C.initSensors()
}

func startAccelerometer(samplesPerSec int) error {
	C.startAccelerometer(C.int(samplesPerSec))
	return nil // TODO(jbd): Return error if no default sensor found.
}

func pollAccelerometer() (deltaX, deltaY, deltaZ float64) {
	e := C.pollAccelerometer()
	return float64(e.x), float64(e.y), float64(e.z)
}

func stopAccelerometer() error {
	C.destroyAccelerometer()
	return nil
}
