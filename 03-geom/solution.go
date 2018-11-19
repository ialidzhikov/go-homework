package main

import (
	"math"

	"github.com/fmi/go-homework/geom"
)

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
	abc := NewTriangle(q.a, q.b, q.c)
	adc := NewTriangle(q.a, q.d, q.c)

	return abc.Intersect(ray) || adc.Intersect(ray)
}

func (s Sphere) Intersect(ray geom.Ray) bool {
	v := subtract(ray.Origin, s.origin)
	b := 2 * dot(v, ray.Direction)
	c := dot(v, v) - s.r*s.r
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
