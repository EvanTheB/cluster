package cluster

import (
	"fmt"
)

type particle struct {
	pos, vel, acc vector3
	invmass       float64
	forceAccum    vector3
}

func (this *particle) integrate(timeStep float64) {
	// book says add curr accel too. constant fields maybe?
	this.acc = scale(this.forceAccum, this.invmass)
	this.forceAccum.zero()
	this.vel.add(scale(this.acc, timeStep))
	// include accel in integration for more accuracy?
	this.pos.add(scale(this.vel, timeStep))
}

func (p *particle) getMass() float64 {
	return 1 / p.invmass
}

func (p *particle) setMass(mass float64) {
	if mass == 0 {
		panic("Mass set to zero")
	}
	p.invmass = 1 / mass
}

func (p *particle) addForce(force vector3) {
	p.forceAccum.add(force)
}

func (p *particle) String() string {
	return fmt.Sprintf("P:%f\nV:%f\nA:%f\nM:%f\n",
		p.pos,
		p.vel,
		p.acc,
		p.getMass(),
	)
}

func (p *particle) getEnergy() float64 {
	magn := mag(p.vel)
	return 0.5 * p.getMass() * magn * magn
}
