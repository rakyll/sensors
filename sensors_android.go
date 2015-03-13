package sensors

/*
#cgo LDFLAGS: -llog -landroid

#include <stdlib.h>
#include <android/sensor.h>

typedef struct AccelerometerEvent {
  int64_t timestamp;
  float x;
  float y;
  float z;
} AccelerometerEvent;

void startAccelerometer();
AccelerometerEvent* pollAccelerometer();

*/
import "C"
import "unsafe"
import "fmt"

func Start() {
	go func() {
		C.startAccelerometer()
		for {
			ev := C.pollAccelerometer()
			fmt.Println(ev.x, ev.y, ev.z)
			C.free(unsafe.Pointer(ev))
		}
	}()
}
