package main

import (
	"engine/objects"
)

type ObjectFactory struct {
	factory *objects.ObjectFactory
}

func NewObjectFactory(f *objects.ObjectFactory) *ObjectFactory {
	return &ObjectFactory{factory: f}
}
