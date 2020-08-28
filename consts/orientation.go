package consts

import "time"

// Orientation ориентация
type Orientation uint8

const (
	OrientationTop    = 0
	OrientationRight  = 1
	OrientationBottom = 2
	OrientationLeft   = 3

	FrameDuration = 30 * time.Millisecond
)
