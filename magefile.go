//go:build mage
// +build mage

package main

import (
	"context"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Go mg.Namespace
type Docker mg.Namespace

const (
	imageRepo = "sujithshajee/dnsbl"
)

var (
	grun   = sh.RunCmd("go", "run")
	gtest  = sh.RunCmd("go", "test")
	ggen   = sh.RunCmd("go", "generate")
	gbuild = sh.RunCmd("go", "build")

	dbuild = sh.RunCmd("docker", "build")
)

func (Go) Test(ctx context.Context) error {
	return gtest("-v", "./...")
}

func (Go) Generate(ctx context.Context) error {
	return ggen("./...")
}

func (Go) Build(ctx context.Context) error {
	return gbuild("-o", "bin/dnsbl", ".")
}

func (Docker) Build() error {
	return dbuild("-t", imageRepo, ".")
}
