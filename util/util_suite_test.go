package util_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taichi/keigo/ginkgo"

	"testing"
)

func TestCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	Configure()
	RunSpecs(t, "util Suite")
}
