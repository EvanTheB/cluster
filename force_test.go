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
	grav := simpleGravity{scale(Z, -10)}
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

func TestRegistry(t *testing.T) {
	reg := getParticleForceRegistry()

	grav := simpleGravity{scale(Z, -10)}

	parts := make([]*particle, 10)
	forces := make([]forceGenerator, 10)
	for i := 0; i < 10; i++ {
		part := particle{
			Z,
			scale(X, 10),
			ZERO,
			1 / 1,
			ZERO,
		}
		parts[i] = &part

		drag := simpleDrag{0.01 * float64(i), 0.01 * float64(i)}
		forces[i] = &drag

		reg.add(&part, &grav)
		reg.add(&part, &drag)
	}
	// Add 10 particles with drag and gravity(shared)
	if len(reg.byForceGenerator) != 11 ||
		len(reg.byParticle) != 10 {
		t.Error("Add failed")
	}

	for i := 0; i < 8; i++ {
		reg.remove(parts[i], forces[i])
	}
	// delete 8 of the particles' drag force
	if len(reg.byForceGenerator) != 3 ||
		len(reg.byParticle) != 10 {
		t.Error("Remove failed")
	}

	reg.removeAllParticle(parts[9])
	// totally remove particle 9 
	if len(reg.byForceGenerator) != 2 ||
		len(reg.byParticle) != 9 {
		t.Error("Remove all Particle failed")
	}

	reg.remove(parts[0], &grav)
	if len(reg.byForceGenerator) != 2 ||
		len(reg.byParticle) != 8 {
		t.Error("Remove last force failed")
	}

	reg.removeAllForce(&grav)
	if len(reg.byForceGenerator) != 1 ||
		len(reg.byParticle) != 1 {
		t.Error("Remove All Force failed")
	}
	// fmt.Println(reg)
}
