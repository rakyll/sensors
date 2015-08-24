// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin
// +build arm arm64

#import <CoreMotion/CoreMotion.h>

void* GoIOS_createManager() {
  return [[CMMotionManager alloc] init];
}

int GoIOS_startAccelerometer(void* m) {
  [((CMMotionManager*)m) startAccelerometerUpdates];
  return 0;
}

int GoIOS_stopAccelerometer(void* m) {
  [((CMMotionManager*)m) stopAccelerometerUpdates];
  return 0;
}

void GoIOS_readAccelerometer(void* m, float* v) {
  CMAccelerometerData* data = ((CMMotionManager*)m).accelerometerData;
  v[0] = (float)data.timestamp;
  v[1] = data.acceleration.x;
  v[2] = data.acceleration.y;
  v[3] = data.acceleration.z;
}

void GoIOS_destroyManager(void* m) {
  [((CMMotionManager*)m) release];
}
