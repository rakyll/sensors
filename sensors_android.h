// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef SENSORS_ANDROID_H
#define SENSORS_ANDROID_H

// Wrapping ASensorEvent, because cgo doesn't support unions.
typedef struct SensorEvent {
  int64_t timestamp;
  float vals[3];
} SensorEvent;

void android_initSensors();
ASensorEventQueue* android_startSensorQueue(int looperId, int type, int32_t usec);
float** android_readSensorQueue(int looperId, ASensorEventQueue* q, int n);
void android_destroySensorQueue(ASensorEventQueue* q);

#endif
