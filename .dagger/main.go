// A generated module for GolangDagger functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/golang-dagger/internal/dagger"
)

type GolangDagger struct{}

// BuildAndTest builds and tests the Go application in the provided source directory
func (m *GolangDagger) BuildAndTest(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.BaseEnv(source).
		WithExec([]string{"go", "build", "."}).
		WithExec([]string{"go", "test", ".", "-v"}).
		Stdout(ctx)
}

// Returns a container that sets up the base environment for the Go application
func (m *GolangDagger) BaseEnv(source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.24.3-alpine3.21").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"go", "get", "-v", "./..."})
}
