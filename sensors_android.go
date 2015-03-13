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
import "unsafe"
import "fmt"

func Start() {
	go func() {
		C.startAccelerometer()
		for {
			ev := C.pollAccelerometer()
			fmt.Println(ev.x, ev.y, ev.z)
			C.free(unsafe.Pointer(ev))
		}
	}()
}
