package core_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RYOT", func() {
	It("should work normally", func() {
		cmd := &RYOT{Term: 2, SwitchTo: On}
		Expect(cmd.String()).To(Equal("RYOT -t 2 TurnOn"))
	})
	It("should work with states", func() {
		cmd := &RYOT{2, TurnOff, 10, 20}
		Expect(cmd.String()).To(Equal("RYOT -t 2 TurnOff -w 10 -t 20"))
	})
})
