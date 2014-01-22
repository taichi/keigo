package core_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RLY", func() {
	It("should work normally", func() {
		cmd := &RLY1{SwitchTo: On}
		Expect(cmd.String()).To(Equal("RLY1 Enable"))
	})
	It("should work normally", func() {
		cmd := &RLY2{SwitchTo: Disable, Wait: 13}
		Expect(cmd.String()).To(Equal("RLY2 Disable -w 13"))
	})
	It("should work normally", func() {
		cmd := &RLY4{SwitchTo: Blink, Wait: 13, Time: 19}
		Expect(cmd.String()).To(Equal("RLY4 Blink -w 13 -t 19"))
	})
})
