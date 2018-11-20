package main

import (
	"math"

	"github.com/fmi/go-homework/geom"
)

const EPSILON float64 = 0.00000001

type Triangle struct {
	A, B, C geom.Vector
}

func NewTriangle(a, b, c geom.Vector) Triangle {
	return Triangle{
		A: a,
		B: b,
		C: c,
	}
}

func (t Triangle) Intersect(ray geom.Ray) bool {
	u := subtract(t.B, t.A)
	v := subtract(t.C, t.A)

	n := crossProduct(u, v)

	direction := subtract(ray.Direction, ray.Origin)
	b := dot(n, direction)

	if math.Abs(b) < EPSILON {
		return false
	}

	w0 := subtract(ray.Origin, t.A)
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
	w := subtract(i, t.A)
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

type Quad struct {
	A, B, C, D geom.Vector
}

func NewQuad(a, b, c, d geom.Vector) Quad {
	return Quad{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func (q Quad) Intersect(ray geom.Ray) bool {
	abc := NewTriangle(q.A, q.B, q.C)
	adc := NewTriangle(q.A, q.D, q.C)

	return abc.Intersect(ray) || adc.Intersect(ray)
}

type Sphere struct {
	Origin geom.Vector
	R      float64
}

func NewSphere(origin geom.Vector, r float64) Sphere {
	return Sphere{
		Origin: origin,
		R:      r,
	}
}

func (s Sphere) Intersect(ray geom.Ray) bool {
	v := subtract(ray.Origin, s.Origin)
	b := 2 * dot(v, ray.Direction)
	c := dot(v, v) - s.R*s.R
	discriminant := b*b - 4*c

	if discriminant < 0 {
		return false
	}

	tMinus := (-b - math.Sqrt(discriminant)) / 2.0
	tPlus := (-b + math.Sqrt(discriminant)) / 2.0

	if tMinus < 0 && tPlus < 0 {
		return false
	}

	return true
}

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

func dot(v, p geom.Vector) float64 {
	return v.X*p.X + v.Y*p.Y + v.Z*p.Z
}
