package cluster

type forceGenerator interface {
	updateForce(part *particle, timeStep float64)
}

type particleForceRegistry struct {
	byParticle       map[*particle][]*forceGenerator
	byForceGenerator map[*forceGenerator][]*particle
}

func (reg *particleForceRegistry) add(part *particle, force *forceGenerator) {
	if reg.byParticle[part] != nil {
		reg.byParticle[part] = append(reg.byParticle[part], force)
	} else {
		arr := make([]*forceGenerator, 5, 0)
		arr = append(arr, force)
		reg.byParticle[part] = arr
	}
	if reg.byForceGenerator[force] != nil {
		reg.byForceGenerator[force] = append(reg.byForceGenerator[force], part)
	} else {
		arr := make([]*particle, 5, 0)
		arr = append(arr, part)
		reg.byForceGenerator[force] = arr
	}
}

func (reg *particleForceRegistry) remove(part *particle, force *forceGenerator) {
	if reg.byParticle[part] != nil {
		reg.byParticle[part] = append(reg.byParticle[part], force)
	} else {
		arr := make([]*forceGenerator, 5, 0)
		arr = append(arr, force)
		reg.byParticle[part] = arr
	}
	if reg.byForceGenerator[force] != nil {
		reg.byForceGenerator[force] = append(reg.byForceGenerator[force], part)
	} else {
		arr := make([]*particle, 5, 0)
		arr = append(arr, part)
		reg.byForceGenerator[force] = arr
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
