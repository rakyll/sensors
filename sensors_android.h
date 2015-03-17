// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef SENSORS_ANDROID_H
#define SENSORS_ANDROID_H

// Wrapping ASensorEvent, because cgo doesn't support unions.
typedef struct AccelerometerEvent {
  int64_t timestamp;
  float x;
  float y;
  float z;
} AccelerometerEvent;

void initSensors();
void startAccelerometer(int samplesPerSec);
AccelerometerEvent pollAccelerometer();
void destroyAccelerometer();

#endif
