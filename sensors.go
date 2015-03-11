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

func Connectivity() (status int) {
	panic("not yet")
}
