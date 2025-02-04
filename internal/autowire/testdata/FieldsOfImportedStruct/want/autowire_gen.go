// Code generated by Autowire. DO NOT EDIT.

//go:generate go run github.com/dabbertorres/autowire/cmd/autowire

//go:build !wireinject
// +build !wireinject

package main

import (
	"example.com/bar"
	"example.com/baz"
	"example.com/foo"
	"fmt"
)

// Injectors from autowire.go:

func newBazService(config *baz.Config) *baz.Service {
	fooConfig := config.Foo
	service := foo.New(fooConfig)
	barConfig := config.Bar
	barService := bar.New(barConfig, service)
	bazService := &baz.Service{
		Foo: service,
		Bar: barService,
	}
	return bazService
}

// autowire.go:

func main() {
	cfg := &baz.Config{
		Foo: &foo.Config{1},
		Bar: &bar.Config{2},
	}
	svc := newBazService(cfg)
	fmt.Println(svc.String())
}
