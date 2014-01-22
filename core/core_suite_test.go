package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taichi/keigo/ginkgo"
	"testing"
)

func TestKeigo(t *testing.T) {
	RegisterFailHandler(Fail)
	Configure()
	RunSpecs(t, "Keigo Suite")
}
