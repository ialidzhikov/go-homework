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

	// todo - test
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
		// todo
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
