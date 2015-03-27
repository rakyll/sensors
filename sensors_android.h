// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef SENSORS_ANDROID_H
#define SENSORS_ANDROID_H

void android_initSensors();
ASensorEventQueue* android_createQueue();
void android_enableSensor(ASensorEventQueue*, int, int32_t);
void android_disableSensor(ASensorEventQueue*, int);
float* android_readQueue(ASensorEventQueue*, int);
void android_destroyQueue(ASensorEventQueue*);

#endif
