package main

import (
	"testing"

	"github.com/fmi/go-homework/geom"
)

func TestSampleSimpleOperations(t *testing.T) {
	var prim geom.Intersectable

	a, b, c := geom.NewVector(-1, -1, 0), geom.NewVector(1, -1, 0), geom.NewVector(0, 1, 0)
	prim = NewTriangle(a, b, c)
	ray := geom.NewRay(geom.NewVector(0, 0, -1), geom.NewVector(0, 0, 1))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, prim)
	}
}

func TestTriangle(t *testing.T) {
	var prim geom.Intersectable

	a, b, c := geom.NewVector(1, 0, 4), geom.NewVector(4, 0, 4), geom.NewVector(2, 2, 4)
	prim = NewTriangle(a, b, c)
	ray := geom.NewRay(geom.NewVector(0, 0, 0), geom.NewVector(1, 0, 0))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect triangle %#v but it did.", ray, prim)
	}

	a, b, c = geom.NewVector(1, 0, 0), geom.NewVector(0, 1, 0), geom.NewVector(0, 0, 0)
	prim = NewTriangle(a, b, c)
	ray = geom.NewRay(geom.NewVector(0.25, 0.25, 1), geom.NewVector(0, 0, 1))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect triangle %#v but it did.", ray, prim)
	}

	// intersection from a side point should return true
	a, b, c = geom.NewVector(1, 0, 0), geom.NewVector(0, 1, 0), geom.NewVector(0, 0, 0)
	prim = NewTriangle(a, b, c)
	ray = geom.NewRay(geom.NewVector(0.25, 0.25, 1), geom.NewVector(1, 0, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, prim)
	}

	// intersection with a short ray should return true
	a, b, c = geom.NewVector(-1, -1, 0), geom.NewVector(1, 0, 0), geom.NewVector(0, 1, 0)
	prim = NewTriangle(a, b, c)
	ray = geom.NewRay(geom.NewVector(0.25, 0.25, 1), geom.NewVector(0.25, 0.25, 0.5))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, prim)
	}

	// intersection with a short ray should return false
	a, b, c = geom.NewVector(-1, -1, 0), geom.NewVector(1, 0, 0), geom.NewVector(0, 1, 0)
	prim = NewTriangle(a, b, c)
	ray = geom.NewRay(geom.NewVector(-2, -2, 1), geom.NewVector(0.25, 0.25, 0.5))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect triangle %#v but it did.", ray, prim)
	}
}

func TestQuad(t *testing.T) {
	var prim geom.Intersectable

	a, b, c, d := geom.NewVector(-1, -1, 0), geom.NewVector(-1, 1, 0), geom.NewVector(1, -1, 0), geom.NewVector(1, 1, 0)
	prim = NewQuad(a, b, c, d)
	ray := geom.NewRay(geom.NewVector(-0.5, -0.5, 1), geom.NewVector(0.5, 0.5, -1))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	// intersection with a parallel ray should return false
	a, b, c, d = geom.NewVector(-1, -1, 0), geom.NewVector(-1, 1, 0), geom.NewVector(1, -1, 0), geom.NewVector(1, 1, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.5, -0.5, 1), geom.NewVector(0.5, 0.5, 1))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	// intersection with a short ray should return true
	a, b, c, d = geom.NewVector(-1, -1, 0), geom.NewVector(-1, 1, 0), geom.NewVector(1, -1, 0), geom.NewVector(1, 1, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.5, -0.5, 1), geom.NewVector(-0.5, -0.5, 0.5))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	// intersection with a short ray should return false
	a, b, c, d = geom.NewVector(-1, -1, 0), geom.NewVector(-1, 1, 0), geom.NewVector(1, -1, 0), geom.NewVector(1, 1, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.5, -0.5, 1), geom.NewVector(-1, -1, 0.5))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	// intersection with a intersection on side should return true
	a, b, c, d = geom.NewVector(-1, -1, 0), geom.NewVector(-1, 1, 0), geom.NewVector(1, -1, 0), geom.NewVector(1, 1, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -0.75, 1), geom.NewVector(-1, -0.75, -1))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	// test with concave quad should return false
	a, b, c, d = geom.NewVector(5, 2.55, 0), geom.NewVector(0.73, 1.31, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(1.75, 3.19, 3.24), geom.NewVector(1.56, 3.25, 0))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	// test with concave quad should return true
	a, b, c, d = geom.NewVector(5, 2.55, 0), geom.NewVector(0.73, 1.31, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.62, -0.23, 1.28), geom.NewVector(-0.56, -0.32, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(5, 2.55, 0), geom.NewVector(0.73, 1.31, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.62, -0.23, 1.28), geom.NewVector(-1.21, 4.49, -0.49))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	a, b, c, d = geom.NewVector(5, 2.55, 0), geom.NewVector(0.73, 1.31, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.62, -0.23, 1.28), geom.NewVector(0.35, 4.5, -0.49))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	a, b, c, d = geom.NewVector(5, 2.55, 0), geom.NewVector(0.73, 1.31, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.62, -0.23, 1.28), geom.NewVector(4.34, -0.1, -0.68))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	a, b, c, d = geom.NewVector(5, 2.55, 0), geom.NewVector(0.73, 1.31, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.62, -0.23, 1.28), geom.NewVector(0.26, 0.65, -0.68))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	a, b, c, d = geom.NewVector(5, 2.55, 0), geom.NewVector(0.73, 1.31, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.62, -0.23, 1.28), geom.NewVector(0.26, 0.65, 1))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(1.93, 3.77, 0), geom.NewVector(2.59, -3.82, 0), geom.NewVector(-1.49, 5.98, 0), geom.NewVector(-1.98, -2.86, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-0.62, -0.23, 1.28), geom.NewVector(0.26, 0.65, 0.67))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-2.46, -3.14, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(3.09, 1.61, -2), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(-1, -1, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did not.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-2.46, -3.14, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(3.09, 1.61, -2), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(-3.53, -4.31, 1))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-2.46, -3.14, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(3.09, 1.61, -2), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(-0.65, 2.99, 1))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-2.46, -3.14, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(-1.26, -1.44, 0.87), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(-0.5, -1.52, 1))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-2.46, -3.14, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(-1.26, -1.44, 0.87), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(0.79, -3.71, 0.24))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-2.46, -3.14, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(-1.26, -1.44, 0.87), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(1.36, -3.42, -0.8))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-2.46, -3.14, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(-1.26, -1.44, 0.87), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(0.91, -3.6, -0.8))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-3.68, -0.65, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(-1.26, -1.44, 0.87), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(0.91, -3.6, -0.8))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}

	a, b, c, d = geom.NewVector(-3.68, -0.65, 2), geom.NewVector(-2.53, 2.82, 0), geom.NewVector(-1.26, -1.44, 0.87), geom.NewVector(2.36, -4.68, 0)
	prim = NewQuad(a, b, c, d)
	ray = geom.NewRay(geom.NewVector(-1, -2, 4), geom.NewVector(1.39, -1.81, -0.8))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect quad %#v but it did.", ray, prim)
	}
}

func TestSphere(t *testing.T) {
	var prim geom.Intersectable

	prim = NewSphere(geom.NewVector(0, 0, 0), 1)
	ray := geom.NewRay(geom.NewVector(0, 4, 0), geom.NewVector(0, -1, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(0, 0, 0), 1)
	ray = geom.NewRay(geom.NewVector(2, -2, 0), geom.NewVector(2, 2, 0))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect sphere %#v but it did.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(0, 0, 0), 1)
	ray = geom.NewRay(geom.NewVector(1, -2, 0), geom.NewVector(1, 2, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(0, 0, 0), 1)
	ray = geom.NewRay(geom.NewVector(1, -1, 0), geom.NewVector(-4, 4, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(0, 0, 0), 1)
	ray = geom.NewRay(geom.NewVector(2.3, -2.3, 0), geom.NewVector(-7.92, 7.65, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(0, 0, 0), 1.5)
	ray = geom.NewRay(geom.NewVector(1, -2, 0), geom.NewVector(1, 2, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(0, 0, 0), 1.5)
	ray = geom.NewRay(geom.NewVector(1, -2, 0), geom.NewVector(1, -4, 0))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect sphere %#v but it did.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(0, 0, 0), 1.5)
	ray = geom.NewRay(geom.NewVector(1, -2, 0), geom.NewVector(5, -4, 2))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect sphere %#v but it did.", ray, prim)
	}

	// test sphere with short ray
	prim = NewSphere(geom.NewVector(0, 0, 0), 1.5)
	ray = geom.NewRay(geom.NewVector(1, -2, 0), geom.NewVector(1, -1.75, 0))

	if !prim.Intersect(ray) {
		// FIXME
		//t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(-4.48, -2.96, 0), geom.NewVector(-7.45, 2.22, 0))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect sphere %#v but it did.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(-4, -3, 0), geom.NewVector(-4, -3, 1))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect sphere %#v but it did.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(-4, -3, 0), geom.NewVector(5.81, 1.47, 0.56))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(-0.72, 1.39, 0), geom.NewVector(5.81, 1.47, 0.56))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(-0.72, 1.39, 0), geom.NewVector(6.78, 1.27, -2.19))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(-0.72, 1.39, 0), geom.NewVector(-3.03, -2.19, -2.19))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect sphere %#v but it did.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(3, 3, 0), geom.NewVector(3, -3, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(3, 0, 0), 2)
	ray = geom.NewRay(geom.NewVector(3, 3, 0), geom.NewVector(3, -3, 3))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(1, 0, 0), 3)
	ray = geom.NewRay(geom.NewVector(4.3, -10.65, 0), geom.NewVector(3.53, 5.92, 0))

	if !prim.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect sphere %#v but it did not.", ray, prim)
	}

	prim = NewSphere(geom.NewVector(1, 0, 0), 3)
	ray = geom.NewRay(geom.NewVector(10, -9, 0), geom.NewVector(9, 8, 0))

	if prim.Intersect(ray) {
		t.Errorf("Did not expected ray %#v to intersect sphere %#v but it did.", ray, prim)
	}
}

func TestSampleIntersectableImplementations(t *testing.T) {
	var prim geom.Intersectable

	a, b, c, d := geom.NewVector(-1, -1, 0),
		geom.NewVector(1, -1, 0),
		geom.NewVector(0, 1, 0),
		geom.NewVector(1, 1, 0)

	prim = NewTriangle(a, b, c)
	prim = NewQuad(a, b, c, d)
	prim = NewSphere(a, 5)

	_ = prim
}
