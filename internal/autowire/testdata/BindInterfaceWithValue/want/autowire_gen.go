// Code generated by Autowire. DO NOT EDIT.

//go:generate go run github.com/dabbertorres/autowire/cmd/autowire

//go:build !wireinject
// +build !wireinject

package main

import (
	"io"
	"os"
)

// Injectors from autowire.go:

func inject() io.Writer {
	file := _wireFileValue
	return file
}

var (
	_wireFileValue = os.Stdout
)
