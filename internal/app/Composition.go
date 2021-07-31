package app

import "example/internal/example"

type Composition struct {
	ExampleController example.Controller
}

func NewComposition() *Composition {
	exampleController := example.NewController()

	return &Composition{
		ExampleController: exampleController,
	}
}
