package cluster

import (
	"fmt"
)

type forceGenerator interface {
	updateForce(part *particle, timeStep float64)
}

type particleForceRegistry struct {
	byParticle       map[*particle][]forceGenerator
	byForceGenerator map[forceGenerator][]*particle
}

func getParticleForceRegistry() particleForceRegistry {
	return particleForceRegistry{
		make(map[*particle][]forceGenerator),
		make(map[forceGenerator][]*particle),
	}
}

func (reg *particleForceRegistry) add(part *particle, force forceGenerator) {
	if reg.byParticle[part] != nil {
		reg.byParticle[part] = append(reg.byParticle[part], force)
	} else {
		arr := make([]forceGenerator, 0, 5)
		arr = append(arr, force)
		reg.byParticle[part] = arr
	}
	if reg.byForceGenerator[force] != nil {
		reg.byForceGenerator[force] = append(reg.byForceGenerator[force], part)
	} else {
		arr := make([]*particle, 0, 5)
		arr = append(arr, part)
		reg.byForceGenerator[force] = arr
	}
}

func (reg *particleForceRegistry) remove(part *particle, force forceGenerator) {
	if arr := reg.byParticle[part]; arr != nil {
		for i := 0; i < len(arr); i++ {
			if arr[i] == force {
				reg.byParticle[part] = append(arr[0:i], arr[i+1:len(arr)]...)
			}
		}
		if len(reg.byParticle[part]) == 0 {
			delete(reg.byParticle, part)
		}
	} else {
		// the particle was not found or the array was nil...
		fmt.Printf("didnt find the particle: %p\n", part)
	}
	if arr := reg.byForceGenerator[force]; arr != nil {
		for i := 0; i < len(arr); i++ {
			if arr[i] == part {
				reg.byForceGenerator[force] = append(arr[0:i], arr[i+1:len(arr)]...)
			}
		}
		if len(reg.byForceGenerator[force]) == 0 {
			delete(reg.byForceGenerator, force)
		}
	} else {
		// the force was not found or the array was nil...
		fmt.Printf("didnt find the force: %p\n", force)
	}
}

func (reg *particleForceRegistry) removeAllParticle(part *particle) {
	if reg.byParticle[part] != nil {
		arr := make([]forceGenerator, len(reg.byParticle[part]))
		copy(arr, reg.byParticle[part])
		for i := 0; i < len(arr); i++ {
			reg.remove(part, arr[i])
		}
	} else {
		// the particle was not found or the array was nil...
		fmt.Printf("didnt find the particle: %p\n", part)
	}
}

func (reg *particleForceRegistry) removeAllForce(force forceGenerator) {
	if reg.byForceGenerator[force] != nil {
		arr := make([]*particle, len(reg.byForceGenerator[force]))
		copy(arr, reg.byForceGenerator[force])
		for i := 0; i < len(arr); i++ {
			reg.remove(arr[i], force)
		}
	} else {
		// the force was not found or the array was nil...
		fmt.Printf("didnt find the force: %p\n", force)
	}
}

type simpleGravity struct {
	gravity vector3
}

func (grav *simpleGravity) updateForce(part *particle, timeStep float64) {
	part.addForce(scale(grav.gravity, part.getMass()))
}

type simpleDrag struct {
	k1, k2 float64
}

func (drag *simpleDrag) updateForce(part *particle, timeStep float64) {
	velMag := mag(part.vel)
	dragForce := velMag*drag.k1 + velMag*velMag*drag.k2
	part.addForce(scale(norm(part.vel), -1*dragForce))
}
