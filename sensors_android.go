package sensors

/*
#cgo LDFLAGS: -llog -landroid

#include <android/log.h>

int start() {
  start_a();
}
*/
import "C"

func Start() {
	go C.start()
}
