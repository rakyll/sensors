package sensors

type Accelerometer struct {
	// TODO(jbd): add criteria
}

func (a *Accelerometer) Start() error {
	panic("not yet implemented")
}

func (a *Accelerometer) HandleFunc(fn func(deltaX, deltaY, deltaZ float64)) {
	panic("not yet implemented")
}

func (a *Accelerometer) Stop() error {
	panic("not yet implemented")
}

type Gyroscope struct {
	// TODO(jbd): add criteria
}

func (g *Gyroscope) Start() error {
	panic("not yet implemented")
}

func (g *Gyroscope) HandleFunc(fn func(roll, pitch, yaw float64)) {
	panic("not implemented")
}

func (g *Gyroscope) Stop() error {
	panic("not yet implemented")
}

// Type of the network that is currently in use.
const (
	TypeWiFi = iota
	TypeMobile
	TypeOther
)

// Connectivity status.
const (
	StatusConnecting = iota
	StatusConnected
	StatusDisconnecting
	StatusUnknown
)

// Connectivity returns the type and the status of the network that is
// currently in use.
func Connectivity() (typ int, status int) {
	panic("not yet")
}
