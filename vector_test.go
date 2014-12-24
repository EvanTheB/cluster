package cluster

import (
	"fmt"
	"testing"
)

func TestVec(t *testing.T) {
	ONE := vector3{1, 1, 1}
	X := vector3{1, 0, 0}
	Y := vector3{0, 1, 0}
	// Z := vector3{0, 0, 1}
	fmt.Println(ONE)
	fmt.Println(scale(ONE, 0.5))
	fmt.Println(add(ONE, X))
	fmt.Println(sub(ONE, X))
	fmt.Println(to(X, Y))
	fmt.Println(dot(X, X))
	fmt.Println(dot(X, Y))
	fmt.Println(cross(X, X))
	fmt.Println(cross(X, Y))
	fmt.Println(mag(scale(ONE, 10)))
	fmt.Println(norm(scale(ONE, 0.1)))
}
