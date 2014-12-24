package cluster

import (
	"fmt"
	"testing"
)

func TestPart(t *testing.T) {
	part := particle{
		ONE,
		ZERO,
		ZERO,
		1 / 1,
		ZERO,
	}
	for i := 0; i < 40; i++ {
		// SHM
		fmt.Println("Step:", i)
		fmt.Println(&part)
		part.addForce(to(part.pos, ZERO))
		part.integrate(0.1)
	}
}
