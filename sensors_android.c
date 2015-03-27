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

ASensorManager* manager = NULL;

// TODO(jbd): looper per each Manager.
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
  return ASensorManager_createEventQueue(manager, looper, LOOPER_ID_ACCELEROMETER, NULL, NULL);
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

int android_readQueue(ASensorEventQueue* q, int n, float* dest) {
  int id;
  int events;
  ASensorEvent event;
  // TODO(jbd): Timeout if pollAll blocks longer than it should.
  int i = 0;
  while (i < n && (id = ALooper_pollAll(-1, NULL, &events, NULL)) >= 0) {
    if (id == LOOPER_ID_ACCELEROMETER) {
      ASensorEvent event;
      if(ASensorEventQueue_getEvents(q, &event, 1)) {
        // TODO(jbd): Handle event type.
        dest[i] = (float)event.type;
        dest[i+1] = (float)event.timestamp;
        dest[i+2] = event.vector.x;
        dest[i+3] = event.vector.y;
        dest[i+4] = event.vector.z;
        i += 5;
      }
    }
  }
  return i;
}

void android_destroyQueue(ASensorEventQueue* q){
  ASensorManager_destroyEventQueue(manager, q);
}