// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdlib.h>
#include <jni.h>

#include <android/log.h>
#include <android/sensor.h>

#include "sensors_android.h"

#define LOG_INFO(...) ((void)__android_log_print(ANDROID_LOG_INFO, "Go/Sensors", __VA_ARGS__))

#define LOOPER_ID_ACCELEROMETER 1

int stopping = 0; // poll event queues until stopping is set.

ASensorManager* manager = NULL;

ALooper* looper = NULL;

void android_initSensors() {
  manager = ASensorManager_getInstance();
}

ASensorEventQueue* android_createQueue() {
  if (looper == NULL) {
    looper = ALooper_forThread();
  }
  if (looper == NULL) {
    looper = ALooper_prepare(ALOOPER_PREPARE_ALLOW_NON_CALLBACKS);
  }
  return ASensorManager_createEventQueue(manager, looper, 0, NULL, NULL);
}

void android_enableSensor(ASensorEventQueue* q, int s, int32_t usec) {
  const ASensor* sensor = ASensorManager_getDefaultSensor(manager, s);
  ASensorEventQueue_enableSensor(q, sensor);
  ASensorEventQueue_setEventRate(q, sensor, usec);
}

void android_disableSensor(ASensorEventQueue* q, int s) {
  const ASensor* sensor = ASensorManager_getDefaultSensor(manager, s);
  ASensorEventQueue_disableSensor(q, sensor);
}

float** android_readQueue(ASensorEventQueue* q, int n) {
  int id;
  int events;
  ASensorEvent event;
  // TODO(jbd): Timeout if pollAll blocks longer than it should.
  float** dest = (float**)malloc(sizeof(float) * 4 * n);
  int i = 0;
  while (i < n && (id = ALooper_pollAll(-1, NULL, &events, NULL)) >= 0) {
    ASensorEvent event;
    if(ASensorEventQueue_getEvents(q, &event, 1)) {
      dest[i][0] = event.timestamp;
      dest[i][1] = event.acceleration.x;
      dest[i][2] = event.acceleration.y;
      dest[i][3] = event.acceleration.z;
    }
    i++;
  }
  return dest;
}
