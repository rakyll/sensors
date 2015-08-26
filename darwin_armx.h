// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin
// +build arm arm64

#ifndef GO_SENSORS_IOS_H
#define GO_SENSORS_IOS_H

#import <CoreMotion/CoreMotion.h>

void* GoIOS_createManager();

int GoIOS_startAccelerometer(void* m);
int GoIOS_stopAccelerometer(void* m);
void GoIOS_readAccelerometer(void* m, float* vectors);

void GoIOS_destroyManager(void* m);

#endif
