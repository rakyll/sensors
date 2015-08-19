// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin
// +build arm arm64

#import <CoreMotion/CoreMotion.h>

void* GoIOS_createManager() {
  return [[CMMotionManager alloc] init];
}

void GoIOS_destroyManager(void* manager) {
  [((CMMotionManager*)manager) release];
}
