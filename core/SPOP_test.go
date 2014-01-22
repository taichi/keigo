package core_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("SPOP", func() {
	It("should work with padding", func() {
		cmd := &SPOP{On, 7, 3}
		Expect(cmd.String()).To(Equal("SPOP 10710300"))
	})
	It("should work normally", func() {
		cmd := &SPOP{SwitchTo: Enable, No: 20}
		Expect(cmd.String()).To(Equal("SPOP 12000000"))
	})
	It("should work simply", func() {
		cmd := &SPOP{}
		Expect(cmd.String()).To(Equal("SPOP"))
	})
	It("should work stop", func() {
		cmd := &SPOP{SwitchTo: Off}
		Expect(cmd.String()).To(Equal("SPOP 00000000"))
	})
})
