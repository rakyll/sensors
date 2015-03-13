// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdlib.h>
#include <jni.h>

#include <android/log.h>
#include <android/sensor.h>

#define LOG_INFO(...) ((void)__android_log_print(ANDROID_LOG_INFO, "Go/Sensors", __VA_ARGS__))

#define LOOPER_ID_ACCELEROMETER 1

int stopping = 0; // poll event queues until stopping is set.

ASensorManager* manager = NULL;

ALooper* aLooper = NULL;
ASensor* aSensor = NULL;
ASensorEventQueue* aEventQueue = NULL;

JNIEXPORT void JNICALL Java_com_Threading_ThreadActivity_stop(JNIEnv *env, jclass clazz) {
  stopping = 1;
}

void start_a() {
  aLooper = ALooper_forThread();
  if (aLooper == NULL) {
    aLooper = ALooper_prepare(ALOOPER_PREPARE_ALLOW_NON_CALLBACKS);
  }
  if (manager == NULL) {
    manager = ASensorManager_getInstance();
  }
  if (aSensor == NULL) {
    aSensor = ASensorManager_getDefaultSensor(manager, ASENSOR_TYPE_ACCELEROMETER);
  }
  aEventQueue = ASensorManager_createEventQueue(manager, aLooper, LOOPER_ID_ACCELEROMETER, NULL, NULL);
  ASensorEventQueue_enableSensor(aEventQueue, aSensor);
  ASensorEventQueue_setEventRate(aEventQueue, aSensor, (1000L/100)*1000);

  int id;  // Identifier.
  int events;
  ASensorEvent event;
  while ((id = ALooper_pollAll(-1, NULL, &events, NULL) >= 0)) {
    if (id == LOOPER_ID_ACCELEROMETER) {
      ASensorEvent event;
      while (ASensorEventQueue_getEvents(aEventQueue, &event, 1) > 0 && !stopping) {
        LOG_INFO("=======================%f %f %f", event.acceleration.x, event.acceleration.y, event.acceleration.z);
      }
    }
     if (stopping) {
      ASensorManager_destroyEventQueue(manager, aEventQueue);
      ALooper_release(aLooper);
    }
  }
}
