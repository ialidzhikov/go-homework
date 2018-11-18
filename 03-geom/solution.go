package main

import "github.com/fmi/go-homework/geom"
import "fmt"
import "math"

const EPSILON float64 = 0.00000001

func add(v, p geom.Vector) geom.Vector {
	return geom.NewVector(v.X+p.X, v.Y+p.Y, v.Z+p.Z)
}

func subtract(v, p geom.Vector) geom.Vector {
	return geom.NewVector(v.X-p.X, v.Y-p.Y, v.Z-p.Z)
}

func scalarProduct(s float64, v geom.Vector) geom.Vector {
	return geom.NewVector(s*v.X, s*v.Y, s*v.Z)
}

func crossProduct(v, p geom.Vector) geom.Vector {
	x := v.Y*p.Z - v.Z*p.Y
	y := v.Z*p.X - v.X*p.Z
	z := v.X*p.Y - v.Y*p.X

	return geom.NewVector(x, y, z)
}

func squaredDistance(v, p geom.Vector) float64 {
	x := p.X - v.X
	y := p.Y - v.Y
	z := p.Z - v.Z

	return x*x + y*y + z*z
}

func dot(v, p geom.Vector) float64 {
	return v.X*p.X + v.Y*p.Y + v.Z*p.Z
}

type Triangle struct {
	a, b, c geom.Vector
}

type Quad struct {
	a, b, c, d geom.Vector
}

type Sphere struct {
	origin geom.Vector
	r      float64
}

func (t Triangle) Intersect(ray geom.Ray) bool {
	u := subtract(t.b, t.a)
	v := subtract(t.c, t.a)

	n := crossProduct(u, v)

	direction := subtract(ray.Direction, ray.Origin)
	b := dot(n, direction)

	if math.Abs(b) < EPSILON {
		return false
	}

	w0 := subtract(ray.Origin, t.a)
	a := dot(n, w0) * -1
	r := a / b
	if r < 0.0 {
		return false
	}

	// Intersection point
	i := add(ray.Origin, scalarProduct(r, direction))

	uu := dot(u, u)
	uv := dot(u, v)
	vv := dot(v, v)
	w := subtract(i, t.a)
	wu := dot(w, u)
	wv := dot(w, v)
	d := uv*uv - uu*vv

	s := (uv*wv - vv*wu) / d
	if s < 0.0 || s > 1.0 {
		return false
	}

	p := (uv*wu - uu*wv) / d
	if p < 0.0 || (s+p) > 1.0 {
		return false
	}

	return true
}

func (q Quad) Intersect(ray geom.Ray) bool {
	// todo
	abc := NewTriangle(q.a, q.b, q.c)
	bcd := NewTriangle(q.b, q.c, q.d)

	return abc.Intersect(ray) || bcd.Intersect(ray)
}

func (s Sphere) Intersect(ray geom.Ray) bool {
	tmp := subtract(s.origin, ray.Origin)
	len := dot(ray.Direction, tmp)
	if len < 0.0 {
		return false
	}

	tmp = add(ray.Origin, scalarProduct(len, ray.Direction))
	dSquare := squaredDistance(s.origin, tmp)
	rSquare := s.r * s.r

	return dSquare <= rSquare
}

func NewTriangle(a, b, c geom.Vector) Triangle {
	return Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func NewQuad(a, b, c, d geom.Vector) Quad {
	return Quad{
		a: a,
		b: b,
		c: c,
		d: d,
	}
}

func NewSphere(origin geom.Vector, r float64) Sphere {
	return Sphere{
		origin: origin,
		r:      r,
	}
}

func main() {
	var prim geom.Intersectable

	a, b, c := geom.NewVector(-1, -1, 0), geom.NewVector(1, -1, 0), geom.NewVector(0, 1, 0)
	prim = NewTriangle(a, b, c)
	ray := geom.NewRay(geom.NewVector(0, 0, -1), geom.NewVector(0, 0, 1))

	if prim.Intersect(ray) {
		fmt.Println("Correct intersection")
	} else {
		fmt.Println("Expected intersection")
	}
}
