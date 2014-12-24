package cluster

import (
	"fmt"
	"testing"
)

func TestForce(t *testing.T) {
	part := particle{
		Z,
		scale(X, 10),
		ZERO,
		1 / 1,
		ZERO,
	}
	grav := simpleGravity{vector3{0, 0, -9.807}}
	drag := simpleDrag{0.1, 0.1}
	for i := 0; i < 40; i++ {
		// SHM
		fmt.Println("Step:", i)
		fmt.Println(&part)
		grav.updateForce(&part, 0.1)
		drag.updateForce(&part, 0.1)
		part.integrate(0.1)
	}
}
