// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdlib.h>
#include <jni.h>

#include <android/log.h>
#include <android/sensor.h>

#include "sensors_android.h"

#define LOG_INFO(...) ((void)__android_log_print(ANDROID_LOG_INFO, "Go/Sensors", __VA_ARGS__))

#define SAMPLES_PER_SEC_ACCELEROMETER 50

#define LOOPER_ID_ACCELEROMETER 1

int stopping = 0; // poll event queues until stopping is set.

ASensorManager* manager = NULL;

ALooper* aLooper = NULL;
ALooper* gLooper = NULL;
ALooper* mLooper = NULL;

ASensorEventQueue* aEventQueue = NULL;
ASensorEventQueue* gEventQueue = NULL;
ASensorEventQueue* mEventQueue = NULL;

JNIEXPORT void JNICALL Java_com_Threading_ThreadActivity_stop(JNIEnv *env, jclass clazz) {
  stopping = 1;
}

void initSensors() {
  manager = ASensorManager_getInstance();
  aLooper = ALooper_forThread();
  if (aLooper == NULL) {
    aLooper = ALooper_prepare(ALOOPER_PREPARE_ALLOW_NON_CALLBACKS);
  }
  gLooper = ALooper_prepare(ALOOPER_PREPARE_ALLOW_NON_CALLBACKS);
  if (gLooper == NULL) {
    gLooper = ALooper_prepare(ALOOPER_PREPARE_ALLOW_NON_CALLBACKS);
  }
  mLooper = ALooper_prepare(ALOOPER_PREPARE_ALLOW_NON_CALLBACKS);
  if (mLooper == NULL) {
    mLooper = ALooper_prepare(ALOOPER_PREPARE_ALLOW_NON_CALLBACKS);
  }
}

void startAccelerometer() {
  const ASensor* sensor = ASensorManager_getDefaultSensor(manager, ASENSOR_TYPE_ACCELEROMETER);
  aEventQueue = ASensorManager_createEventQueue(manager, aLooper, LOOPER_ID_ACCELEROMETER, NULL, NULL);
  ASensorEventQueue_enableSensor(aEventQueue, sensor);
  ASensorEventQueue_setEventRate(aEventQueue, sensor, (1000L/SAMPLES_PER_SEC_ACCELEROMETER)*1000);
}

void destroyAccelerometer() {
  ASensorManager_destroyEventQueue(manager, aEventQueue);
  ALooper_release(aLooper);
}

AccelerometerEvent* pollAccelerometer() {
  // TODO(jbd): Investigate how much of a bottleneck
  // to poll from the event queue one by one.
  int id;
  int events;
  ASensorEvent event;
  AccelerometerEvent* e = NULL;
  if ((id = ALooper_pollOnce(-1, NULL, &events, NULL)) >= 0) {
     if (id == LOOPER_ID_ACCELEROMETER) {
      ASensorEvent event;
      if (ASensorEventQueue_getEvents(aEventQueue, &event, 1) > 0 && !stopping) {
        e = (AccelerometerEvent*)malloc(sizeof(struct AccelerometerEvent));
        e->timestamp = event.timestamp;
        e->x = event.acceleration.x;
        e->y = event.acceleration.y;
        e->z = event.acceleration.z;
        return e;
      }
    }
     if (stopping) {
      destroyAccelerometer();
    }
  }
  return e;
}
