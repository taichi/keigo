package util_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taichi/keigo/util"
)

var _ = Describe("Functions", func() {
	Describe("ToOnceFn", func() {
		var (
			n  int
			fn func()
		)
		BeforeEach(func() {
			n = 0
			fn = ToOnceFn(func() {
				n++
			})
		})
		It("should execute once", func() {
			fn()
			fn()
			fn()
			Expect(n).To(Equal(1))
		})
	})
})
